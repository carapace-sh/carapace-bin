package os

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/condition"
)

// ActionDisplays completes x displays
//
//	:0
//	:1
func ActionDisplays() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "darwin":
			if !condition.Executable("w")(c) {
				return carapace.ActionValues()
			}
			return carapace.ActionExecCommand("w", "-hsu")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				r := regexp.MustCompile(`(?:/usr/lib/Xorg|X11.app|Xquartz) (?P<display>:[0-9]+(\.[0-9]+)?)`)
				for line := range strings.SplitSeq(string(output), "\n") {
					if r.MatchString(line) {
						vals = append(vals, r.FindStringSubmatch(line)[1])
					}
				}
				return carapace.ActionValues(vals...)
			})
		case "windows":
			return carapace.ActionValues() // Windows doesn't use X display identifiers
		default:
			return carapace.ActionExecCommand("w", "-hsu")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				r := regexp.MustCompile("/usr/lib/Xorg (?P<display>:[0-9]+)")
				for line := range strings.SplitSeq(string(output), "\n") {
					if r.MatchString(line) {
						vals = append(vals, r.FindStringSubmatch(line)[1])
					}
				}
				return carapace.ActionValues(vals...)
			})
		}
	}).Tag("displays")
}
