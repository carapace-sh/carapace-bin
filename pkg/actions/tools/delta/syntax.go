package delta

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionSyntaxThemes completes syntax themes
//
//	1337 (dark)
//	Catppuccin-mocha (dark)
func ActionSyntaxThemes() carapace.Action {
	return carapace.ActionExecCommand("delta", "--list-syntax-themes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if splitted := strings.SplitN(line, "\t", 2); len(splitted) == 2 {
				if t := splitted[0]; t == "dark" {
					vals = append(vals, splitted[1], splitted[0], style.Blue)
				} else {
					vals = append(vals, splitted[1], splitted[0], style.Default)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
