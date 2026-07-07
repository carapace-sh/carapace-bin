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
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
