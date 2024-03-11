package brew

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAllFormulae completes all formulae
func ActionAllFormulae() carapace.Action {
	return carapace.ActionExecCommand("brew", "formulae")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("formulae")
}
