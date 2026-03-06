package wt

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionBranches completes branches
func ActionBranches() carapace.Action {
	return carapace.Batch(
		git.ActionLocalBranches(),
		git.ActionRemoteBranches(""),
	).ToA() // TODO verify
}
