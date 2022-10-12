package gum

import "github.com/rsteube/carapace"

// ActionAlignments completes ActionAlignments
//
//	left
//	center
func ActionAlignments() carapace.Action {
	return carapace.ActionValues("left", "center", "right", "bottom", "middle", "top")
}
