package tmux

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionWindows completes windows
//
//	0 (elvish- (1 panes) [106x28] [layout c97e,106x28,0,0,1] @1)
//	1 (elvish* (1 panes) [106x28] [layout c97f,106x28,0,0,2] @2 (active))
func ActionWindows() carapace.Action {
	return carapace.ActionExecCommand("tmux", "list-windows")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.SplitN(line, ": ", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
