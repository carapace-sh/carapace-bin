package action

import (
	"github.com/rsteube/carapace"
)

func ActionYesNo() carapace.Action {
	return carapace.ActionValues("yes", "no")
}
