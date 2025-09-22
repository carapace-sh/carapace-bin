package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionStashes completes stash names
//
//	stash@{0} (WIP on branchA: ABCDEF1 last commit msg)
//	stash@{1} (WIP on branchB: ABCDEF2 last commit msg)
func ActionStashes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "stash", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, ": ", 2)
				vals = append(vals, splitted[0], splitted[1])
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Stash)
		})
	}).Tag("stashes").UidF(Uid("stash")).QueryF(Uid("stash"))
}
