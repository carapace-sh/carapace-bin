package tmux

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBuffers completes buffers
//
//	0 (buffer content...)
//	1 (buffer content...)
func ActionBuffers() carapace.Action {
	return carapace.ActionExecCommand("tmux", "list-buffers")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			if splitted := strings.SplitN(line, ": ", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
