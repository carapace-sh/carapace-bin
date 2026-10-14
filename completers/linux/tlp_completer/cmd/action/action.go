package action

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBatteries completes battery names (e.g. BAT0, BAT1).
func ActionBatteries() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := filepath.Glob("/sys/class/power_supply/BAT*")
		if err != nil || len(entries) == 0 {
			return carapace.ActionValues("BAT0", "BAT1")
		}

		batteries := make([]string, 0, len(entries))
		for _, e := range entries {
			batteries = append(batteries, filepath.Base(e))
		}
		sort.Strings(batteries)
		return carapace.ActionValues(batteries...)
	})
}

// ActionConfigParams completes `CONFIG_PARAM=` overrides
// appended after `--` (see tlp(8) CONFIGURATION section).
func ActionConfigParams() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		params := make(map[string]struct{})
		for _, file := range []string{"/usr/share/tlp/defaults.conf", "/etc/tlp.conf"} {
			content, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			for _, line := range strings.Split(string(content), "\n") {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				if name, _, found := strings.Cut(line, "="); found {
					name = strings.TrimSpace(name)
					if name != "" && !strings.Contains(name, " ") {
						params[name] = struct{}{}
					}
				}
			}
		}

		names := make([]string, 0, len(params))
		for name := range params {
			names = append(names, name)
		}
		sort.Strings(names)
		return carapace.ActionValues(names...).Suffix("=")
	})
}
