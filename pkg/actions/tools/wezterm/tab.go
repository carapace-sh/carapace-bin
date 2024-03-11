package wezterm

import (
	"strconv"

	"github.com/carapace-sh/carapace"
)

// ActionTabs completes tabs
//
//	0
//	1
func ActionTabs() carapace.Action {
	return actionPanes(func(panes []pane) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			vals = append(vals, strconv.Itoa(p.TabID), p.TabTitle)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
