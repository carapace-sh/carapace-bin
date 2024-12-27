package ghostty

import "github.com/carapace-sh/carapace"

// ActionResizeOverlayModes completes resize overlay modes
//
//	always (Always show resize overlays)
//	never (Never show resize overlays)
func ActionResizeOverlayModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "Always show resize overlays",
		"never", "Never show resize overlays",
		"after-first", "Show up if the surface is resized after creation",
	).Tag("resize overlay modes")
}

// ActionResizeOverlayPositions completes resize overlay positions
//
//	center
//	top-left
func ActionResizeOverlayPositions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"center",
		"top-left",
		"top-center",
		"top-right",
		"bottom-left",
		"bottom-center",
		"bottom-right",
	).Tag("resize overlay positions")
}
