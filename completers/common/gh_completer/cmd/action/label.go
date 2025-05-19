package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type label struct {
	Name        string
	Description string
	Color       string
}

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return gh.ActionLabels(gh.RepoOpts{Host: repo.RepoHost(), Owner: repo.RepoOwner(), Name: repo.RepoName()})
	})
}
