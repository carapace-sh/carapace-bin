package gh

import (
	"fmt"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/util"
)

type environment struct {
	Name      string
	UpdatedAt time.Time `json:"updated_at"`
}

type environmentsQuery struct {
	Environments []environment
}

// ActionEnvironments completes environments
//
//	production (updated 3 days ago)
//	staging (updated 2 hours ago)
func ActionEnvironments(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult environmentsQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/environments", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, environment := range queryResult.Environments {
					vals = append(vals, environment.Name, fmt.Sprintf("updated %v", util.FuzzyAgo(time.Since(environment.UpdatedAt))))
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("environments")
}
