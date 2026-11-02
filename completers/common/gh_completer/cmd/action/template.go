package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionIssueTemplates completes issue templates
func ActionIssueTemplates(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionIssueTemplates(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionPullRequestTemplates completes pull request templates
func ActionPullRequestTemplates(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestTemplates(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
