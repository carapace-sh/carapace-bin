package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionDiscussions completes discussions
func ActionDiscussions(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionDiscussions(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionDiscussionCategories completes discussion categories
func ActionDiscussionCategories(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionDiscussionCategories(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
