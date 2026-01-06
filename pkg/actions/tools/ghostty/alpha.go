package ghostty

import "github.com/carapace-sh/carapace"

// ActionAlphaBlendings completes alpha blendings
//
//	native (Perform alpha blending in the native color space for the OS)
//	linear (Perform alpha blending in linear space)
func ActionAlphaBlendings() carapace.Action {
	return carapace.ActionValuesDescribed(
		"native", "Perform alpha blending in the native color space for the OS",
		"linear", "Perform alpha blending in linear space",
		"linear-corrected", "Same as `linear`, but with a correction step applied",
	)
}
