package scoop

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAliases completes alias names with their command as description
//
//	rm (scoop uninstall $args[0])
//	upgrade (scoop update *)
func ActionAliases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		root := loadConfig().scoopDir()
		if root == "" {
			return carapace.ActionValues()
		}
		matches, err := filepath.Glob(filepath.Join(root, "shims", "scoop-*.ps1"))
		if err != nil {
			return carapace.ActionValues()
		}

		sort.Strings(matches)
		vals := make([]string, 0, len(matches)*2)
		for _, match := range matches {
			name := strings.TrimSuffix(strings.TrimPrefix(filepath.Base(match), "scoop-"), ".ps1")
			vals = append(vals, name, aliasCommand(match))
		}
		return carapace.ActionValuesDescribed(vals...).Tag("aliases")
	})
}

func aliasCommand(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	if len(lines) < 2 {
		return ""
	}
	return strings.TrimSpace(lines[1])
}
