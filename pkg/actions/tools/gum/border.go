package gum

import "github.com/carapace-sh/carapace"

// ActionBorders completes borders
//
//	none
//	hidden
func ActionBorders() carapace.Action {
	return carapace.ActionValuesDescribed(
		"double", "╔═",
		"hidden", "",
		"none", "",
		"normal", "┌─",
		"rounded", "╭─",
		"thick", "┏━",
	)
}
