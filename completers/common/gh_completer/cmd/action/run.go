package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

type RunOpts struct {
	All        bool
	InProgress bool
	Failed     bool
	Successful bool
}

// ActionWorkflowRuns completes workflow runs
func ActionWorkflowRuns(cmd *cobra.Command, opts RunOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionWorkflowRuns(gh.RunOpts{
		Host:       repo.Host,
		Owner:      repo.Owner,
		Name:       repo.Name,
		All:        opts.All,
		InProgress: opts.InProgress,
		Failed:     opts.Failed,
		Successful: opts.Successful,
	})
}

// ActionRunFields completes run fields
func ActionRunFields() carapace.Action {
	return gh.ActionRunFields()
}

// ActionWorkflowJobs completes jobs of a workflow run
func ActionWorkflowJobs(cmd *cobra.Command, runId string, _ RunOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionWorkflowJobs(gh.JobsOpts{
		Host:  repo.Host,
		Owner: repo.Owner,
		Name:  repo.Name,
		RunId: runId,
	})
}
