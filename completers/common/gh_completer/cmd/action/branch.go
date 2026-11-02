package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionBranches completes branches
func ActionBranches(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionBranches(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionBranchCommits completes commits of a branch
func ActionBranchCommits(cmd *cobra.Command, branch string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionBranchCommits(gh.BranchOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Branch: branch,
	})
}
