package ghostty

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionThemes completes themes
//
//	Atom
//	Ayu
func ActionThemes() carapace.Action {
	return carapace.ActionExecCommand("ghostty", "+list-themes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			vals = append(vals, strings.Split(line, " (")[0])
		}
		return carapace.ActionValues(vals...)
	}).Tag("themes")
}
