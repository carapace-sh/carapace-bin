package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionBackgroundImageFits completes background image fits
//
//	contain (scale to the largest size that can still be contained)
//	cover (scale to the smallest size that can completely cover the terminal)
func ActionBackgroundImageFits() carapace.Action {
	return carapace.ActionValuesDescribed(
		"contain", "scale to the largest size that can still be contained",
		"cover", "scale to the smallest size that can completely cover the terminal",
		"stretch", "stretch the background image to the full size of the terminal",
		"none", "don't scale the background image",
	).StyleF(style.ForKeyword).
		Tag("background image fits")
}

// ActionBackgroundImagePositions completes background image positions
//
//	top-left
//	top-center
func ActionBackgroundImagePositions() carapace.Action {
	return carapace.ActionValues(
		"top-left",
		"top-center",
		"top-right",
		"center-left",
		"center",
		"center-right",
		"bottom-left",
		"bottom-center",
		"bottom-right",
	).Tag("background image positions")
}
