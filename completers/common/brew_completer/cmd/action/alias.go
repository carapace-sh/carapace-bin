package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAliases completes brew aliases
func ActionAliases() carapace.Action {
	return carapace.ActionExecCommand("brew", "alias")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("aliases")
}
