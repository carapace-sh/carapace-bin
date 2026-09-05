package bazel

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionCommands completes bazel commands
//
//	analyze-profile (Analyzes build profile data)
//	aquery (Analyzes the given targets and the query output.)
func ActionCommands() carapace.Action {
	return carapace.ActionExecCommand("bazel", "help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "Usage:") || strings.HasPrefix(line, "Getting") {
				continue
			}
			fields := strings.Fields(line)
			if len(fields) > 0 {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("commands")
}