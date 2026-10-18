package tmux

import "github.com/carapace-sh/carapace"

// ActionLayouts completes layout names
//
//	even-horizontal
//	even-vertical
func ActionLayouts() carapace.Action {
	return carapace.ActionValues(
		"even-horizontal",
		"even-vertical",
		"main-horizontal",
		"main-vertical",
		"tiled",
	)
}
