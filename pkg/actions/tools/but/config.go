package but

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

// ActionConfigNames completes config names
//
//	email
//	editor
func ActionConfigNames() carapace.Action {
	return carapace.ActionValues(
		"name",
		"email",
		"editor",
	)
}

// ActionConfigValues completes config values
func ActionConfigValues(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionMap{
			"editor": bridge.ActionCarapaceBin().Split(),
		}[name]
	})
}
