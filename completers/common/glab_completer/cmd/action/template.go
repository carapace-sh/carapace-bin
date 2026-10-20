package action

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionTemplates(templateDir string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if filepath.IsAbs(c.Value) || strings.ContainsAny(c.Value, `/\`) {
			return carapace.ActionFiles()
		}

		root, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			return carapace.ActionFiles()
		}

		entries, err := os.ReadDir(filepath.Join(strings.TrimSpace(string(root)), ".gitlab", templateDir))
		if err != nil {
			return carapace.ActionFiles()
		}

		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			names = append(names, strings.TrimSuffix(entry.Name(), ".md"))
		}
		return carapace.Batch(
			carapace.ActionValues(names...),
			carapace.ActionFiles(),
		).Invoke(c).Merge().ToA()
	})
}
