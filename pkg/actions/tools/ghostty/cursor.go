package ghostty

import "github.com/carapace-sh/carapace"

// ActionCursorStyles completes cursor styles
//
//	block
//	bar
func ActionCursorStyles() carapace.Action {
	return carapace.ActionValues(
		"block",
		"bar",
		"underline",
		"block_hollow",
	).Tag("cursor styles")
}
