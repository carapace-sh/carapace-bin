package os

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionSessionIds completes session ids
//
//	0 (root)
//	1 (root)
func ActionSessionIds() carapace.Action {
	return carapace.ActionExecCommand("ps", "-A", "-o", "user,sess")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		ids := make(map[string]string)

		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) == 2 {
				ids[fields[1]] = fields[0]
			}
		}

		vals := make([]string, 0)
		for key, value := range ids {
			vals = append(vals, key, value)
		}

		return carapace.ActionValuesDescribed(vals...)
	}).Tag("session ids")
}
