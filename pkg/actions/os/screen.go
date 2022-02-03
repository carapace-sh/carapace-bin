package os

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionScreens completes sreens
//   eDP1
//   HDMI1
func ActionScreens(connected bool) carapace.Action {
	return carapace.ActionExecCommand("xrandr")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		var r *regexp.Regexp
		if connected {
			r = regexp.MustCompile(`^(?P<screen>[^ ]+) connected .*`)
		} else {
			r = regexp.MustCompile(`^(?P<screen>[^ ]+) disconnected .*`)
		}

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
