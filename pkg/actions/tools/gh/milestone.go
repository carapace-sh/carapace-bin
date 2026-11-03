package gh

import (
	"fmt"
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
)

type milestone struct {
	Title       string
	Description string
}

type milestonesQuery struct {
	Data struct {
		Repository struct {
			Milestones struct {
				Nodes []milestone
			}
		}
	}
}

// ActionMilestones completes milestones
//
//	v1.0 (first stable release)
//	v2.0 (API breaking changes)
func ActionMilestones(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult milestonesQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ milestones(first: 100) { nodes { title, description } } }`, &queryResult, func() carapace.Action {
			milestones := queryResult.Data.Repository.Milestones.Nodes
			vals := make([]string, len(milestones)*2)
			for index, milestone := range milestones {
				vals[index*2] = milestone.Title
				vals[index*2+1] = milestone.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Cache(24*time.Hour, opts.cacheKey())
	}).Tag("milestones")
}

type milestonesRestQuery struct {
	Number int
	Title  string
	State  string
}

// ActionMilestoneNumbers completes milestone numbers
//
//	3 (v1.0, active)
//	4 (v2.0, closed)
func ActionMilestoneNumbers(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult []milestonesRestQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/milestones?state=all&per_page=100", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0, len(queryResult)*2)
				for _, milestone := range queryResult {
					vals = append(vals, strconv.Itoa(milestone.Number), fmt.Sprintf("%v (%v)", milestone.Title, milestone.State))
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("milestone numbers")
}
