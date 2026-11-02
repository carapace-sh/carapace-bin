package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type IssueOpts struct {
	Closed bool
	Open   bool
}

// ActionIssues completes issues
func ActionIssues(cmd *cobra.Command, opts IssueOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionIssues(gh.IssueOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Closed: opts.Closed,
		Open:   opts.Open,
	})
}

// ActionIssueFields completes issue fields
func ActionIssueFields() carapace.Action {
	return gh.ActionIssueFields()
}

// ActionIssueLabels completes labels of an issue
func ActionIssueLabels(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionIssueLabels(gh.IssueOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionIssueAssignees completes assignees of an issue
func ActionIssueAssignees(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionIssueAssignees(gh.IssueOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionIssueTypes completes issue types
func ActionIssueTypes(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionIssueTypes(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionPinnedIssues completes pinned issues
func ActionPinnedIssues(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPinnedIssues(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
