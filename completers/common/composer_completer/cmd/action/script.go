package action

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionScripts completes the scripts defined in the composer.json
func ActionScripts(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadProject(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		descriptions := p.object("scripts-descriptions")
		vals := make([]string, 0)
		for name := range p.object("scripts") {
			description, _ := descriptions[name].(string)
			vals = append(vals, name, oneLine(description))
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("scripts")
}

// ActionBinaries completes the binaries of the bin-dir and the ones defined in the composer.json
func ActionBinaries(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		p, err := loadProject(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		dir, err := workingDir(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vendor, err := vendorDir(c, cmd, p)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		binDir := c.Getenv("COMPOSER_BIN_DIR")
		if binDir == "" {
			binDir, _ = p.config()["bin-dir"].(string)
		}
		if binDir == "" {
			binDir = "{$vendor-dir}/bin"
		}
		binDir = strings.ReplaceAll(binDir, "{$vendor-dir}", filepath.ToSlash(vendor))
		if !filepath.IsAbs(binDir) {
			binDir = filepath.Join(dir, binDir)
		}

		names := make(map[string]bool)
		if entries, err := os.ReadDir(binDir); err == nil {
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				if name, ok := strings.CutSuffix(entry.Name(), ".bat"); ok && names[name] {
					continue // windows proxy of a binary that is already listed
				}
				names[entry.Name()] = true
			}
		}
		for _, bin := range p.strings("bin") {
			names[filepath.Base(bin)] = true
		}

		if len(names) == 0 {
			return carapace.ActionMessage("no binaries found")
		}

		vals := make([]string, 0, len(names))
		for name := range names {
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...)
	}).Tag("binaries")
}
