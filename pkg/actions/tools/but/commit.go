package but

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionCommits completes commits
//
//	36ae34b (some commit)
//	e1b2490 (another commit)
func ActionCommits() carapace.Action {
	return actionStatus(false, func(status butStatus) carapace.Action {
		vals := make([]string, 0)
		for _, stack := range status.Stacks {
			for _, branch := range stack.Branches {
				for _, commit := range branch.Commits {
					vals = append(vals, commit.CommitID[:7], commit.Message, styles.Git.Commit)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("commits").UidF(git.Uid("ref")) // TODO custom uid
}
