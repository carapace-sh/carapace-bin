package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionContents completes contents
func ActionContents(cmd *cobra.Command, path string, branch string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionContents(gh.ContentOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Branch: branch,
		Path:   path,
	})
}
