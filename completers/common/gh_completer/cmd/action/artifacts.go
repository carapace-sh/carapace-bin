package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionWorkflowArtifactNames completes artifact names
func ActionWorkflowArtifactNames(cmd *cobra.Command, runId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionWorkflowArtifactNames(gh.ArtifactOpts{
		Host:  repo.Host,
		Owner: repo.Owner,
		Name:  repo.Name,
		RunId: runId,
	})
}

// ActionWorkflowArtifactIds completes artifact ids
func ActionWorkflowArtifactIds(cmd *cobra.Command, runId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionWorkflowArtifactIds(gh.ArtifactOpts{
		Host:  repo.Host,
		Owner: repo.Owner,
		Name:  repo.Name,
		RunId: runId,
	})
}
