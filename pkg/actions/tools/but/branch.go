package but

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionLocalBranches completes local branches
//
//	mybranch
//	another
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("but", "branch", "list", "--local")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		branches := make([]string, 0)

		for _, line := range lines { // TODO branch list has no json output yet
			if fields := strings.Fields(line); len(fields) == 2 {
				branches = append(branches, fields[1])
			}
		}
		// but branches are git branches and lack context, so just use those from git
		return git.ActionLocalBranches().Retain(branches...)
	})
}
