package action

import "github.com/carapace-sh/carapace"

func ActionVcsTypes() carapace.Action {
	return carapace.ActionValues("bitbucket", "github")
}
