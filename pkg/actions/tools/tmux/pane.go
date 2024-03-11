package tmux

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPanes completes panes
func ActionPanes() carapace.Action {
	return carapace.ActionExecCommand("tmux", "list-panes")(func(output []byte) carapace.Action {
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
