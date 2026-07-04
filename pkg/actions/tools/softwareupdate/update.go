package softwareupdate

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionUpdates completes available software update item names
//
//	macOS Big Sur 11.7.1-1234567890
//	Safari 15.6-9876543210
func ActionUpdates() carapace.Action {
	return carapace.ActionExecCommand("softwareupdate", "--list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "* ") {
				// Lines like: "* Label: macOS Big Sur 11.7.1-1234567890, Title: macOS Big Sur 11.7.1, Recommended: YES, Action: restart,"
				if prefix, rest, ok := strings.Cut(line, "Label: "); ok && prefix == "* " {
					if label, _, ok := strings.Cut(rest, ","); ok {
						vals = append(vals, strings.TrimSpace(label))
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("updates").Uid("softwareupdate", "updates")
}
