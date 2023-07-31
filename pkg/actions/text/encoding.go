package text

import (
	"github.com/rsteube/carapace"
)

// ActionEncodings completes encodings
//
//	ASCII
//	UTF-8
func ActionEncodings() carapace.Action {
	// TODO incomplete - move to different package than text?
	return carapace.ActionValues(
		"ASCII",
		"UTF-8",
		"UTF-16",
	)
}
