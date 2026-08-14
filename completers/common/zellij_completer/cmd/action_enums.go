package cmd

import "github.com/carapace-sh/carapace"

func actionDirections() carapace.Action {
	return carapace.ActionValues("left", "down", "up", "right")
}

func actionResizes() carapace.Action {
	return carapace.ActionValues("increase", "decrease")
}
