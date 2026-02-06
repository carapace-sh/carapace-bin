package scrcpy

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionApps completes apps
//
//	com.github.android (GitHub)
//	com.github.rsteube.t4 (T4 Launcher)
func ActionApps() carapace.Action {
	return carapace.ActionExecCommand("scrcpy", "--list-apps")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ (?P<type>[-*]) (?P<name>.+)    +(?P<id>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				switch matches[1] {
				case "-":
					vals = append(vals, matches[3], matches[2], style.Default)
				case "*":
					vals = append(vals, matches[3], matches[2], style.Yellow)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("apps")
}
