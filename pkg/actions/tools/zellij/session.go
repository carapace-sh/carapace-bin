package zellij

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSessions completes active zellij sessions
//
//	main
//	dev
func ActionSessions() carapace.Action {
	return carapace.ActionExecCommand("zellij", "list-sessions", "--short", "--no-formatting")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		var vals []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "No active zellij sessions") {
				continue
			}
			vals = append(vals, line)
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no active sessions")
		}
		return carapace.ActionValues(vals...)
	}).Tag("sessions")
}
