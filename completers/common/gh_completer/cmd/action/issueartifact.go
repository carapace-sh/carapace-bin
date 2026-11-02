package action

import (
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionIssueArtifacts completes issue artifacts
func ActionIssueArtifacts(cmd *cobra.Command, issueNumber string) carapace.Action {
	repo := repoOpts(cmd)
	number, err := strconv.Atoi(issueNumber)
	if err != nil {
		return carapace.ActionValues()
	}
	return gh.ActionIssueArtifacts(gh.IssueArtifactOpts{
		Host:        repo.Host,
		Owner:       repo.Owner,
		Name:        repo.Name,
		IssueNumber: number,
	})
}

// ActionIssueArtifactVersions completes issue artifact versions
func ActionIssueArtifactVersions(cmd *cobra.Command, issueNumber, artifactNumber string) carapace.Action {
	repo := repoOpts(cmd)
	number, err := strconv.Atoi(issueNumber)
	if err != nil {
		return carapace.ActionValues()
	}
	artifact, err := strconv.Atoi(artifactNumber)
	if err != nil {
		return carapace.ActionValues()
	}
	return gh.ActionIssueArtifactVersions(gh.IssueArtifactVersionOpts{
		Host:           repo.Host,
		Owner:          repo.Owner,
		Name:           repo.Name,
		IssueNumber:    number,
		ArtifactNumber: artifact,
	})
}
