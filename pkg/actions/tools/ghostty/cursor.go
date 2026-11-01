package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCursorStyles completes cursor styles
//
//	block
//	bar
func ActionCursorStyles() carapace.Action {
	return carapace.ActionValues(
		"bar",
		"block",
		"block_hollow",
		"underline",
	).Tag("cursor styles").Uid("ghostty", "cursor-style")
}

// ActionCursorStyleBlinks completes cursor style blink modes
//
//	true (The cursor blinks by default)
//	false (The cursor does not blink)
func ActionCursorStyleBlinks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"false", "The cursor does not blink",
		"true", "The cursor blinks by default",
	).StyleF(style.ForKeyword).Tag("cursor style blink modes").Uid("ghostty", "cursor-style-blink")
}
