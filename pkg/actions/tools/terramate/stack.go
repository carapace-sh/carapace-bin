package terramate

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStacks completes stacks.
//
//	first
//	second
func ActionStacks() carapace.Action {
	return carapace.ActionExecCommand("terramate", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
