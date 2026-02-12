package atuin

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAliases completes aliases
//
//	b (bat)
//	g (git)
func ActionAliases() carapace.Action {
	return carapace.ActionExecCommand("atuin", "dotfiles", "alias", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if name, value, ok := strings.Cut(line, "="); ok {
				vals = append(vals, name, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("aliases")
}
