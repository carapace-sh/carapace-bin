package action

import "github.com/carapace-sh/carapace"

func ActionKinds() carapace.Action {
	return carapace.ActionValues("home", "page", "section", "taxonomy", "term").UniqueList(",")
}
