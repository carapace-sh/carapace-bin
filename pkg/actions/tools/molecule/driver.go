package molecule

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDrivers completes molecule drivers
func ActionDrivers() carapace.Action {
	return carapace.ActionExecCommand("molecule", "drivers", "-f", "plain")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) +(?P<descrioption>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("ansible drivers")
}
