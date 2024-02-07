package brew

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionAllCasks completes all casks
func ActionAllCasks() carapace.Action {
	return carapace.ActionExecCommand("brew", "casks")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("casks")
}
