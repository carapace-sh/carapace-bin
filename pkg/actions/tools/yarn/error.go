package yarn

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionErrorCodes completes error codes
//
//	YN0000 (UNNAMED)
//	YN0001 (EXCEPTION)
func ActionErrorCodes() carapace.Action {
	return carapace.ActionExecCommand("yarn", "explain")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			if fields := strings.Fields(line); len(fields) == 3 {
				vals = append(vals, strings.TrimSuffix(fields[1], ":"), fields[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
