package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

func ActionDiscussions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return gh.ActionDiscussions(gh.RepoOpts{Host: repo.RepoHost(), Owner: repo.RepoOwner(), Name: repo.RepoName()})
	})
}

func ActionDiscussionCategories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return gh.ActionDiscussionCategories(gh.RepoOpts{Host: repo.RepoHost(), Owner: repo.RepoOwner(), Name: repo.RepoName()})
	})
}
