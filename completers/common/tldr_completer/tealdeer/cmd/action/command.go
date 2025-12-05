package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCommands() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tldr", "--list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines...)
		})
	})
}
