package action

import "github.com/rsteube/carapace"

func ActionKinds() carapace.Action {
	return carapace.ActionValues("page", "home", "section", "taxonomy", "term").UniqueList(",")
}
