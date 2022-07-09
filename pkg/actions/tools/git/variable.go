package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionVariables completes variables
//   GIT_EDITOR (nvim)
//   GIT_PAGER (bat --style grid)
func ActionVariables() carapace.Action {
	return carapace.ActionExecCommand("git", "var", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, "=", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
