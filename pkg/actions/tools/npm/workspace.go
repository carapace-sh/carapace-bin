package npm

import (
	"github.com/carapace-sh/carapace"
)

// ActionWorkspaces completes workspaces
func ActionWorkspaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		pj, err := loadPackageJson(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(pj.Workspaces...)
	})
}
