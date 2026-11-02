package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionEnvironments completes environments
func ActionEnvironments(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionEnvironments(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
