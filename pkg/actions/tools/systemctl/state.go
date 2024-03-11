package systemctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStates completes states
//
//	dead
//	running
func ActionStates() carapace.Action {
	// TODO handle different units
	return carapace.ActionExecCommand("systemctl", "--state=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if !strings.HasPrefix(line, "Available ") && line != "" {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
