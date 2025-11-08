package but

import "github.com/carapace-sh/carapace"

// ActionTargets completes targets
//
//	36ae34b (some commit)
//	branch  (some branch)
func ActionTargets() carapace.Action {
	return carapace.Batch(
		ActionCommits(),
		ActionLocalBranches(),
	).ToA()
}
