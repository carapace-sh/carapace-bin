package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

// ActionGlobal completes the arguments of the command given to `composer global` (executed in the global composer directory)
func ActionGlobal() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home, err := composerHome(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return bridge.ActionCarapaceBin("composer", "--working-dir="+home)
	})
}
