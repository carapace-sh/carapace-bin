package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionBody completes issue/pr bodies
func ActionBody(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionBody(gh.BodyOpts(repo))
}
