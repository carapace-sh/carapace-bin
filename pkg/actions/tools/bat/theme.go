package bat

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionThemes completes themes
//
//	DarkNeon
//	TwoDark
func ActionThemes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bat", "--list-themes")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
