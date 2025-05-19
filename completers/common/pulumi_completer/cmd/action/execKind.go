package action

import (
	"github.com/carapace-sh/carapace"
)

func ActionExecKinds() carapace.Action {
	return carapace.ActionValues(
		"auto.local",
		"auto.inline",
		"cli",
	)
}
