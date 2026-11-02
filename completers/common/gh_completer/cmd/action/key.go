package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionDeployKeys completes deploy keys
func ActionDeployKeys(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionDeployKeys(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
