package gh

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionWorkflowEvents completes events triggering a workflow
//
//	create
//	delete
func ActionWorkflowEvents() carapace.Action {
	return carapace.ActionValues(
		"branch_protection_rule",
		"check_run",
		"check_suite",
		"create",
		"delete",
		"deployment",
		"deployment_status",
		"discussion",
		"discussion_comment",
		"fork",
		"gollum",
		"issue_comment",
		"issues",
		"label",
		"merge_group",
		"milestone",
		"page_build",
		"project",
		"project_card",
		"project_column",
		"public",
		"pull_request",
		"pull_request_review",
		"pull_request_review_comment",
		"pull_request_target",
		"push",
		"registry_package",
		"release",
		"repository_dispatch",
		"schedule",
		"status",
		"watch",
		"workflow_call",
		"workflow_dispatch",
		"workflow_run",
	)
}

// ActionWorkflowFields completes workflow fields.
//
//	name
//	state
func ActionWorkflowFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"name",
		"path",
		"state",
	)
}

type workflow struct {
	Id    int
	Name  string
	State string
}

type workflowsQuery struct {
	Workflows []workflow
}

type WorkflowOpts struct {
	Host  string
	Owner string
	Name  string

	Enabled  bool
	Disabled bool

	ById   bool
	ByName bool
}

func (o WorkflowOpts) Default() WorkflowOpts {
	o.Enabled = true
	o.Disabled = true
	o.ByName = true
	return o
}

func (o WorkflowOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

func (o WorkflowOpts) cacheKey() key.Key {
	return key.String(o.Host, o.Owner, o.Name,
		strconv.FormatBool(o.Enabled),
		strconv.FormatBool(o.Disabled),
		strconv.FormatBool(o.ById),
		strconv.FormatBool(o.ByName),
	)
}

// ActionWorkflows completes workflows
//
//	1234567 (build)
//	2345678 (test)
//
//	carapace.ActionWorkflows(gh.WorkflowOpts{Enabled: true})
func ActionWorkflows(opts WorkflowOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult workflowsQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/actions/workflows", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, workflow := range queryResult.Workflows {
					if opts.Enabled && workflow.State == "active" ||
						opts.Disabled && workflow.State != "active" {
						if opts.ById {
							vals = append(vals, strconv.Itoa(workflow.Id), workflow.Name)
						}
						if opts.ByName {
							vals = append(vals, workflow.Name, strconv.Itoa(workflow.Id))
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("workflows").Cache(24*time.Hour, opts.cacheKey())
}
