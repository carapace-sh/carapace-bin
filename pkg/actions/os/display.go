package os

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDisplays completes x displays
//   :0
//   :1
func ActionDisplays() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("w", "-hsu")(func(output []byte) carapace.Action {
			vals := make([]string, 0)
			r := regexp.MustCompile("/usr/lib/Xorg (?P<display>:[0-9]+)")
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					vals = append(vals, r.FindStringSubmatch(line)[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
