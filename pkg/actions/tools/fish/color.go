package fish

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionColorNames completes fish color names
//
//	black (black)
//	red (red)
func ActionColorNames() carapace.Action {
	return carapace.ActionExecCommand("fish", "-c", "set_color --print-colors")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], strings.Join(fields[1:], " "))
				} else {
					vals = append(vals, line)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
