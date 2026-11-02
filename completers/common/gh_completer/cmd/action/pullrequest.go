package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type PullRequestOpts struct {
	Closed    bool
	Merged    bool
	Open      bool
	DraftOnly bool
}

// ActionPullRequests completes pull requests
func ActionPullRequests(cmd *cobra.Command, opts PullRequestOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequests(gh.PullRequestOpts{
		Host:      repo.Host,
		Owner:     repo.Owner,
		Name:      repo.Name,
		Closed:    opts.Closed,
		Merged:    opts.Merged,
		Open:      opts.Open,
		DraftOnly: opts.DraftOnly,
	})
}

// ActionPullRequestFields completes pull request fields
func ActionPullRequestFields() carapace.Action {
	return gh.ActionPullRequestFields()
}

// ActionPullRequestCheckFields completes pull request check fields
func ActionPullRequestCheckFields() carapace.Action {
	return gh.ActionPullRequestCheckFields()
}

// ActionPullRequestLabels completes labels of a pull request
func ActionPullRequestLabels(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestLabels(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestAssignees completes assignees of a pull request
func ActionPullRequestAssignees(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestAssignees(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestReviewers completes reviewers of a pull request
func ActionPullRequestReviewers(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestReviewers(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestCommits completes commits of a pull request
func ActionPullRequestCommits(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestCommits(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}
