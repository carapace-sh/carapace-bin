package mitmproxy

import "github.com/carapace-sh/carapace"

// ActionConsoleLayouts completes console layouts
//
//	horizontal
//	single
func ActionConsoleLayouts() carapace.Action {
	return carapace.ActionValues("horizontal", "single", "vertical")
}
