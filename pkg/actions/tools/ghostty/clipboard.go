package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionClipboardPermissions completes clipboard permissions
//
//	ask (Ask the user for permission each time a program tries to access the clipboard)
//	allow (Allow clipboard access without prompting)
func ActionClipboardPermissions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ask", "Ask the user for permission each time a program tries to access the clipboard",
		"allow", "Allow clipboard access without prompting",
		"deny", "Deny clipboard access",
	).StyleF(style.ForKeyword).Tag("clipboard permissions").Uid("ghostty", "clipboard-permission")
}
