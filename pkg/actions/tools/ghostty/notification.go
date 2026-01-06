package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionNotifications completes app notifications
//
//	clipboard-copy (Show a notification when text is copied to the clipboard)
//	config-reload (Show a notification when the configuration is reloaded)
func ActionNotifications() carapace.Action {
	return carapace.ActionValuesDescribed(
		"clipboard-copy", "Show a notification when text is copied to the clipboard",
		"config-reload", "Show a notification when the configuration is reloaded",
		"no-clipboard-copy", "Do not show a notification when text is copied to the clipboard",
		"no-config-reload", "Do not show a notification when the configuration is reloaded",
		"false", "disable all notifications",
		"true", "will enable all notifications",
	).StyleF(style.ForKeyword)
}
