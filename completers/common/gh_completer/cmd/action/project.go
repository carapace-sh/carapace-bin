package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type ProjectOpts struct {
	Closed bool
	Open   bool
}

// ActionProjects completes legacy repository (classic) projects
func ActionProjects(cmd *cobra.Command, opts ProjectOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionRepositoryProjects(gh.RepositoryProjectOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Closed: opts.Closed,
		Open:   opts.Open,
	})
}
