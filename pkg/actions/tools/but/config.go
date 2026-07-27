package but

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionConfigNames completes config names
//
//	email
//	editor
func ActionConfigNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "Git user name (user.name)",
		"email", "Git user email (user.email)",
		"editor", "Git editor (core.editor)",
	)
}

// ActionConfigValues completes config values
//
//	true
//	false
func ActionConfigValues(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionMap{
			"editor": bridge.ActionCarapaceBin().Split(),
		}[name]
	})
}

// ActionUIConfigNames completes UI config names
//
//	email
//	editor
func ActionUIConfigNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"tui", "Use the interactive TUI for diff by default (but.ui.tui)",
	)
}

// ActionUIConfigValues completes config values
//
//	true
//	false
func ActionUIConfigValues(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionMap{
			"tui": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		}[name]
	})
}
