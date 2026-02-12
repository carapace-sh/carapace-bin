package atuin

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionVariables completes variables
//
//	one (1)
//	two (2)
func ActionVariables() carapace.Action {
	return carapace.ActionExecCommand("atuin", "dotfiles", "var", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if name, value, ok := strings.Cut(strings.TrimPrefix(line, "export "), "="); ok {
				vals = append(vals, name, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("variables")
}
