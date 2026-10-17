package rex

import "github.com/carapace-sh/carapace"

// ActionDirections completes block split directions
//
//	above (place the block above the anchor)
//	right (place the block to the right of the anchor)
func ActionDirections() carapace.Action {
	return carapace.ActionValuesDescribed(
		"above", "place the block above the anchor",
		"below", "place the block below the anchor",
		"left", "place the block to the left of the anchor",
		"right", "place the block to the right of the anchor",
	).Tag("directions")
}
