package pi

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes installed packages
//
//	npm:@foo/pi-tools
//	git:github.com/user/repo
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("pi", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "    ") {
				trimmed := strings.TrimSpace(line)
				if idx := strings.Index(trimmed, " ("); idx > 0 {
					trimmed = trimmed[:idx]
				}
				vals = append(vals, trimmed)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("packages")
}
