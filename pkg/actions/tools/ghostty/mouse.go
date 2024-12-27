package ghostty

import "github.com/carapace-sh/carapace"

// ActionMouseShiftCaptureModes completes mouse shift capture modes
//
//	always (always capture)
//	never (never capture)
func ActionMouseShiftCaptureModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "capture, but allow override",
		"false", "don't capture, but allow override",
		"always", "always capture",
		"never", "never capture",
	).Tag("mouse shift capture modes")
}
