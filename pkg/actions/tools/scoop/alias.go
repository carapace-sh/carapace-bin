package scoop

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// Aliases returns configured aliases with their command as value
//
//	rm (scoop uninstall $args[0])
//	upgrade (scoop update *)
func Aliases() map[string]string {
	cfg := loadConfig()
	root := cfg.scoopDir()
	if root == "" {
		return nil
	}

	configured := cfg.aliases()
	aliases := make(map[string]string, len(configured))
	for name := range configured {
		path := filepath.Join(root, "shims", "scoop-"+name+".ps1")
		if _, err := os.Stat(path); err != nil {
			aliases[name] = "<BROKEN>"
			continue
		}
		aliases[name] = aliasCommand(path)
	}
	return aliases
}

// ActionAliases completes alias names with their command as description
//
//	rm (scoop uninstall $args[0])
//	upgrade (scoop update *)
func ActionAliases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		aliases := Aliases()
		names := make([]string, 0, len(aliases))
		for name := range aliases {
			names = append(names, name)
		}
		sort.Strings(names)

		vals := make([]string, 0, len(names)*2)
		for _, name := range names {
			vals = append(vals, name, aliases[name])
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
