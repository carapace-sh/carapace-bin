package golang

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionTools completes tools
//
//	asm
//	buildid
func ActionTools() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("go", "tool")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
