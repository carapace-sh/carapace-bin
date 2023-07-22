package gum

import "github.com/rsteube/carapace"

// ActionCursorModes completes cursor modes
func ActionCursorModes() carapace.Action {
	return carapace.ActionValues("blink", "hide", "static")
}
