package jj

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionRemotes completes remotes
//
//	carapace (https://github.com/rsteube/carapace)
//	carapace-bin (https://github.com/rsteube/carapace-bin)
func ActionRemotes() carapace.Action {
	return carapace.ActionExecCommand("jj", "git", "remote", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, " ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("remotes")
}
