package jj

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionConflicts completes conflicts
func ActionConflicts(revision string) carapace.Action {
	return carapace.ActionExecCommand("jj", "resolve", "--list", "--revision", revision)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, "    ", 2)...) // TODO verify
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
