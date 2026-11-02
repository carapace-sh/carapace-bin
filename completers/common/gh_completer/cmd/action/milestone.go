package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionMilestones completes milestones
func ActionMilestones(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionMilestones(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
