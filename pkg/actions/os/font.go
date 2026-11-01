package os

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/condition"
)

// ActionFontFamilies completes font family names
//
//	MesloLGSDZ Nerd Font
//	Nimbus Sans
func ActionFontFamilies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch {
		case condition.Executable("fc-list")(c):
			return carapace.ActionExecCommand("fc-list", "--format=%{family}\n")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		case runtime.GOOS == "darwin":
			return carapace.ActionExecCommand("system_profiler", "SPFontsDataType", "-detailLevel", "mini")(func(output []byte) carapace.Action {
				families := make(map[string]bool)
				for line := range strings.SplitSeq(string(output), "\n") {
					line = strings.TrimSpace(line)
					if prefix, value, ok := strings.Cut(line, ":"); ok && strings.TrimSpace(prefix) == "family" {
						name := strings.TrimSpace(value)
						if name != "" {
							families[name] = true
						}
					}
				}
				vals := make([]string, 0, len(families))
				for family := range families {
					vals = append(vals, family)
				}
				return carapace.ActionValues(vals...)
			})
		case runtime.GOOS == "windows":
			return actionFontFamiliesWindows()
		default:
			return carapace.ActionValues()
		}
	}).Tag("font families")
}

func actionFontFamiliesWindows() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		families := make(map[string]bool)
		fontDirs := []string{
			`C:\Windows\Fonts`,
		}
		if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
			fontDirs = append(fontDirs, userProfile+`\AppData\Local\Microsoft\Windows\Fonts`)
		}
		for _, dir := range fontDirs {
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				name := entry.Name()
				// Strip extension and suffixes like "(TrueType)" or "(OpenType)"
				name = strings.TrimSuffix(name, filepath.Ext(name))
				name = strings.TrimSpace(strings.TrimSuffix(name, " (TrueType)"))
				name = strings.TrimSpace(strings.TrimSuffix(name, " (OpenType)"))
				if name != "" {
					families[name] = true
				}
			}
		}
		vals := make([]string, 0, len(families))
		for family := range families {
			vals = append(vals, family)
		}
		return carapace.ActionValues(vals...)
	})
}
