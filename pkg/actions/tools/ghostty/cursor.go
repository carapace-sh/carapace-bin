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
		"block",
		"bar",
		"underline",
		"block_hollow",
	).Tag("cursor styles").Uid("ghostty", "cursor-style")
}

// ActionCursorStyleBlinks completes cursor style blink modes
//
//	true (The cursor blinks by default)
//	false (The cursor does not blink)
func ActionCursorStyleBlinks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "The cursor blinks by default",
		"false", "The cursor does not blink",
		"in", "The cursor blinks only when the terminal is focused",
		"out", "The cursor blinks only when the terminal is unfocused",
	).StyleF(style.ForKeyword).Tag("cursor style blink modes").Uid("ghostty", "cursor-style-blink")
}
