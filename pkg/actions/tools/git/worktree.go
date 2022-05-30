package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionWorktrees completes worktrees
//   /tmp/carapace-spec.git ((bare))
//   /tmp/carapace-spec.git/mexample (a3d9c48 [doc-usage])
func ActionWorktrees() carapace.Action {
	return carapace.ActionExecCommand("git", "worktree", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
