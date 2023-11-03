package jj

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionWorkspaces completes workspaces
//
//	default (qzzmpvmx 2ceef6bf (no description set)
//	another (oxtpukyp 00a745c4 (empty) (no description set))
func ActionWorkspaces() carapace.Action {
	return carapace.ActionExecCommand("jj", "workspace", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, ": ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
