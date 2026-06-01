package gum

import "github.com/carapace-sh/carapace"

// ActionCursorModes completes cursor modes
//
//	blink
//	hide
func ActionCursorModes() carapace.Action {
	return carapace.ActionValues("blink", "hide", "static")
}
