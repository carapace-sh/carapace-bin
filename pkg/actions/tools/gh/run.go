package gh

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace-bin/pkg/util"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/carapace-sh/carapace/pkg/style"
)

type run struct {
	Id         int
	Name       string
	Status     string
	Conclusion string
	HeadBranch string    `json:"head_branch"`
	CreatedAt  time.Time `json:"created_at"`
}

type runsQuery struct {
	WorkflowRuns []run `json:"workflow_runs"`
}

type RunOpts struct {
	Host  string
	Owner string
	Name  string

	All        bool
	InProgress bool
	Failed     bool
	Successful bool
}

func (o RunOpts) Default() RunOpts {
	o.All = true
	return o
}

func (o RunOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

func (o RunOpts) cacheKey() key.Key {
	return key.String(o.Host, o.Owner, o.Name,
		strconv.FormatBool(o.All),
		strconv.FormatBool(o.InProgress),
		strconv.FormatBool(o.Failed),
		strconv.FormatBool(o.Successful),
	)
}

func (o RunOpts) includes(r run) bool {
	switch {
	case o.All:
		return true
	case o.InProgress && r.Status == "in_progress":
		return true
	case o.Failed && r.Status == "completed" && r.Conclusion != "success":
		return true
	case o.Successful && r.Status == "completed" && r.Conclusion == "success":
		return true
	default:
		return false
	}
}

// ActionWorkflowRuns completes workflow runs
//
//	14469905264 (build (master) just now)
func ActionWorkflowRuns(opts RunOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult runsQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/actions/runs", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, run := range queryResult.WorkflowRuns {
					if opts.includes(run) {
						ago := time.Since(run.CreatedAt)
						vals = append(vals, strconv.Itoa(run.Id), fmt.Sprintf("%v (%v) %v", run.Name, run.HeadBranch, util.FuzzyAgo(ago)), styleForRun(run))
					}
				}
				return carapace.ActionStyledValuesDescribed(vals...)
			})
		})
	}).Tag("workflow runs")
}

func styleForRun(run run) string {
	switch run.Status {
	case "in_progress":
		return styles.Gh.JobInProgress
	case "completed":
		if run.Conclusion == "success" {
			return styles.Gh.JobSuccess
		}
		return styles.Gh.JobFailed
	default:
		return style.Default
	}
}

type job struct {
	Id         int
	Name       string
	StartedAt  time.Time `json:"started_at"`
	Status     string
	Conclusion string
}

type jobsQuery struct {
	Jobs []job
}

type JobsOpts struct {
	Host  string
	Owner string
	Name  string

	RunId string
}

func (o JobsOpts) Default() JobsOpts {
	return o
}

func (o JobsOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

// ActionWorkflowJobs completes jobs of a workflow run
//
//	9984259356 (build (1.44.1))
func ActionWorkflowJobs(opts JobsOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult jobsQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/actions/runs/%v/jobs", resolved.Owner, resolved.Name, opts.RunId), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, job := range queryResult.Jobs {
					vals = append(vals, strconv.Itoa(job.Id), job.Name, styleForJob(job))
				}
				return carapace.ActionStyledValuesDescribed(vals...)
			})
		})
	}).Tag("jobs")
}

func styleForJob(job job) string {
	switch job.Status {
	case "in_progress":
		return styles.Gh.JobInProgress
	case "completed":
		if job.Conclusion == "success" {
			return styles.Gh.JobSuccess
		}
		return styles.Gh.JobFailed
	default:
		return style.Default
	}
}

// ActionRunFields completes run fields
//
//	displayTitle
//	status
func ActionRunFields() carapace.Action {
	return carapace.ActionValues(
		"conclusion",
		"createdAt",
		"databaseId",
		"displayTitle",
		"event",
		"headBranch",
		"headSha",
		"name",
		"number",
		"startedAt",
		"status",
		"updatedAt",
		"url",
		"workflowDatabaseId",
		"workflowName",
	)
}
