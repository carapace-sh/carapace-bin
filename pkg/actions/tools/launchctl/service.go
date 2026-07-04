package launchctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionServices completes launchd service labels
//
//	com.apple.something
//	system.log
func ActionServices() carapace.Action {
	return carapace.ActionExecCommand("launchctl", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				label := fields[2]
				if label != "Label" && strings.TrimSpace(label) != "" {
					vals = append(vals, label)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("services").Uid("tools.launchctl", "services")
}
