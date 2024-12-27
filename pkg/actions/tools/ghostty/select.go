package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionCopyOnSelectModes completes copy-on-select modes
//
//	true (prefer to copy to the selection clipboard if supported by the OS)
//	clipboard (always copy text to the selection clipboard)
func ActionCopyOnSelectModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "prefer to copy to the selection clipboard if supported by the OS",
		"clipboard", "always copy text to the selection clipboard",
		"false", "do not automatically copy selected text to the clipboard",
	).StyleF(style.ForKeyword).Tag("copy-on-select modes")
}
