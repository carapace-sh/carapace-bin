package ghostty

import "github.com/carapace-sh/carapace"

// ActionGtkToolbarStyles completes adw toolbar styles
//
//	flat (Top and bottom bars are flat with the terminal window)
//	raised (Top and bottom bars cast a shadow on the terminal area)
func ActionGtkToolbarStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"flat", "Top and bottom bars are flat with the terminal window",
		"raised", "Top and bottom bars cast a shadow on the terminal area",
		"raised-border", "Similar to `raised` but the shadow is replaced with a more subtle border",
	).Tag("gtk toolbar styles")
}
