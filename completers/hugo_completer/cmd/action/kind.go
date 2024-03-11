package action

import "github.com/carapace-sh/carapace"

func ActionKinds() carapace.Action {
	return carapace.ActionValues("page", "home", "section", "taxonomy", "term").UniqueList(",")
}
