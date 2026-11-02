package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionTeams completes teams
func ActionTeams(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionTeams(gh.OwnerOpts{Host: repo.Host, Owner: repo.Owner})
}
