package gum

import "github.com/carapace-sh/carapace"

// ActionCursorModes completes cursor modes
func ActionCursorModes() carapace.Action {
	return carapace.ActionValues("blink", "hide", "static")
}
