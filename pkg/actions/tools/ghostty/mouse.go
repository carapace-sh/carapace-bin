package ghostty

import "github.com/carapace-sh/carapace"

// ActionMouseShiftCaptureModes completes mouse shift capture modes
//
//	always (always capture)
//	never (never capture)
func ActionMouseShiftCaptureModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "always capture",
		"false", "don't capture, but allow override",
		"never", "never capture",
		"true", "capture, but allow override",
	).Tag("mouse shift capture modes").Uid("ghostty", "mouse-shift-capture")
}
