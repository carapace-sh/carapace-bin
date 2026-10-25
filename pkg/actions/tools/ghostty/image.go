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
		"none", "don't scale the background image",
		"stretch", "stretch the background image to the full size of the terminal",
	).StyleF(style.ForKeyword).
		Tag("background image fits").Uid("ghostty", "background-image-fit")
}

// ActionBackgroundImagePositions completes background image positions
//
//	top-left
//	top-center
func ActionBackgroundImagePositions() carapace.Action {
	return carapace.ActionValues(
		"bottom-center",
		"bottom-left",
		"bottom-right",
		"center",
		"center-left",
		"center-right",
		"top-center",
		"top-left",
		"top-right",
	).Tag("background image positions").Uid("ghostty", "background-image-position")
}
