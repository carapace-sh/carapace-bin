package wezterm

import (
	"strconv"

	"github.com/carapace-sh/carapace"
)

// ActionWindows completes windows
//
//	0
//	1
func ActionWindows() carapace.Action {
	return actionPanes(func(panes []pane) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			vals = append(vals, strconv.Itoa(p.WindowID), p.WindowTitle)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
