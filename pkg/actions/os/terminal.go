package os

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionTerminals completes terminals
//
//	pts/5 (user)
//	tty7 (root)
func ActionTerminals() carapace.Action {
	return carapace.ActionExecCommand("ps", "aux")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		terminals := make(map[string]string)

		for _, line := range lines[1:] {
			if fields := strings.Fields(line); len(fields) > 6 && fields[6] != "?" {
				terminals[fields[6]] = fields[0]
			}
		}

		vals := make([]string, 0)
		for key, value := range terminals {
			vals = append(vals, key, value)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("terminals")
}
