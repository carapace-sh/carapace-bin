package action

import "github.com/carapace-sh/carapace"

func ActionLogLevel() carapace.Action {
	return carapace.ActionValues(
		"quiet",
		"info",
		"error",
	)
}
