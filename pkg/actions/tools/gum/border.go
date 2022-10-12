package gum

import "github.com/rsteube/carapace"

// ActionBorders completes borders
//
//	none
//	hidden
func ActionBorders() carapace.Action {
	return carapace.ActionValues("none", "hidden", "normal", "rounded", "thick", "double")
}
