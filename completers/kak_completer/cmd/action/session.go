package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionSessions() carapace.Action {
	return carapace.ActionExecCommand("kak", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
