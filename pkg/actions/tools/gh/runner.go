package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type runner struct {
	Id     int
	Name   string
	Os     string `json:"os"`
	Status string
}

type runnersQuery struct {
	Runners []runner
}

// ActionRepositoryRunners completes self-hosted runners of a repository
//
//	1234567 (cicd-runner (Linux))
func ActionRepositoryRunners(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult runnersQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/actions/runners", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				return runnerAction(queryResult.Runners)
			})
		})
	}).Tag("runners")
}

// ActionOrganizationRunners completes self-hosted runners of an organization
//
//	1234567 (cicd-runner (Linux))
func ActionOrganizationRunners(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult runnersQuery
		return apiV3Action(opts.repo(), fmt.Sprintf("orgs/%v/actions/runners", opts.Owner), &queryResult, func() carapace.Action {
			return runnerAction(queryResult.Runners)
		})
	}).Tag("runners")
}

func runnerAction(runners []runner) carapace.Action {
	vals := make([]string, 0, len(runners)*2)
	for _, runner := range runners {
		vals = append(vals, fmt.Sprintf("%v", runner.Id), fmt.Sprintf("%v (%v)", runner.Name, runner.Os))
	}
	return carapace.ActionValuesDescribed(vals...)
}

type runnerGroup struct {
	Id        int
	Name      string
	IsDefault bool `json:"default"`
}

type runnerGroupsQuery struct {
	RunnerGroups []runnerGroup `json:"runner_groups"`
}

// ActionOrganizationRunnerGroups completes self-hosted runner groups of an organization
//
//	1234567 (Default)
func ActionOrganizationRunnerGroups(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult runnerGroupsQuery
		return apiV3Action(opts.repo(), fmt.Sprintf("orgs/%v/actions/runner-groups", opts.Owner), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult.RunnerGroups)*2)
			for _, group := range queryResult.RunnerGroups {
				vals = append(vals, fmt.Sprintf("%v", group.Id), group.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("runner groups")
}
