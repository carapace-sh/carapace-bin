package gum

import "github.com/carapace-sh/carapace"

// ActionAlignments completes ActionAlignments
//
//	left
//	center
func ActionAlignments() carapace.Action {
	return carapace.ActionValues("bottom", "center", "left", "middle", "right", "top")
}
