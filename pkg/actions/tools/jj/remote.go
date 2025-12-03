package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionRemotes completes remotes
//
//	carapace (https://github.com/carapace-sh/carapace)
//	carapace-bin (https://github.com/carapace-sh/carapace-bin)
func ActionRemotes() carapace.Action {
	return actionExecJJ("git", "remote", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, " ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("remotes")
}
