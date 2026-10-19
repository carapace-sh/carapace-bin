package scoop

import (
	"os"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionCachedApps completes cached app names with their version as description
//
//	bat (0.24.0)
//	git (2.45.0)
func ActionCachedApps() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir := loadConfig().cacheDir()
		if dir == "" {
			return carapace.ActionValues()
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			return carapace.ActionValues()
		}

		versions := make(map[string]string)
		names := make([]string, 0)
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			parts := strings.Split(entry.Name(), "#")
			if len(parts) < 2 || parts[0] == "" {
				continue
			}
			if _, ok := versions[parts[0]]; ok {
				continue
			}
			names = append(names, parts[0])
			versions[parts[0]] = parts[1]
		}

		sort.Strings(names)
		vals := make([]string, 0, len(names)*2)
		for _, name := range names {
			vals = append(vals, name, versions[name])
		}
		return carapace.ActionValuesDescribed(vals...).Tag("cached apps")
	})
}
