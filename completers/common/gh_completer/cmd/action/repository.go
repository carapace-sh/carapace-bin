package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionOwnerRepositories completes owner/repository
func ActionOwnerRepositories(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionOwnerRepositories(gh.HostOpts{Host: repo.Host})
}
