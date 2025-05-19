package action

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace-bin/pkg/util"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type run struct {
	Id         int
	Name       string
	Status     string
	Conclusion string
	HeadBranch string    `json:"head_branch"`
	CreatedAt  time.Time `json:"created_at"`
}

type RunOpts struct {
	All        bool
	InProgress bool
	Failed     bool
	Successful bool
}

type runQuery struct {
	WorkflowRuns []run `json:"workflow_runs"`
}

func ActionWorkflowRuns(cmd *cobra.Command, opts RunOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult runQuery
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/actions/runs`, repo.RepoOwner(), repo.RepoName()), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, run := range queryResult.WorkflowRuns {
				ago := time.Since(run.CreatedAt)
				if opts.All ||
					(opts.InProgress && run.Status == "in_progress") ||
					(opts.Failed && run.Status == "completed" && run.Conclusion != "success") ||
					(opts.Successful && run.Status == "completed" && run.Conclusion == "success") {

					vals = append(vals, strconv.Itoa(run.Id), fmt.Sprintf("%v (%v) %v", run.Name, run.HeadBranch, util.FuzzyAgo(ago)), styleForRun(run))
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func styleForRun(run run) string {
	switch run.Status {
	case "in_progress":
		return styles.Gh.JobInProgress
	case "completed":
		if run.Conclusion == "success" {
			return styles.Gh.JobSuccess
		} else {
			return styles.Gh.JobFailed
		}
	default:
		return style.Default
	}
}

func ActionRunFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"displayTitle",
		"headBranch",
		"headSha",
		"createdAt",
		"updatedAt",
		"startedAt",
		"status",
		"conclusion",
		"event",
		"number",
		"databaseId",
		"workflowDatabaseId",
		"workflowName",
		"url",
	)
}
