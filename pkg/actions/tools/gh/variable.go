package gh

import (
	"fmt"
	"time"

	"github.com/rsteube/carapace"
)

type variable struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EnvironmentOpts struct {
	Host        string
	Owner       string
	Name        string
	Environment string
}

func (o EnvironmentOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner}

}

// ActionEnvironmentVariables completes environment variables
//
//	first (value1)
//	second (value2)
func ActionEnvironmentVariables(opts EnvironmentOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Variables []variable
		}
		return apiV3Action(opts.repo(), fmt.Sprintf("repos/%v/%v/environments/github-pages/variables?per_page=100", opts.Owner, opts.Name), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, v := range queryResult.Variables {
				vals = append(vals, v.Name, v.Value) // TODO updated at
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionRepositoryVariables completes repository variables
//
//	first (value1)
//	second (value2)
func ActionRepositoryVariables(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Variables []variable
		}
		return apiV3Action(opts, fmt.Sprintf("repos/%v/%v/actions/variables?per_page=100", opts.Owner, opts.Name), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, v := range queryResult.Variables {
				vals = append(vals, v.Name, v.Value) // TODO updated at
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionOrganisationVariables completes organisation variables
//
//	first (value1)
//	second (value2)
func ActionOrganisationVariables(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Variables []variable
		}
		return apiV3Action(opts.repo(), fmt.Sprintf("orgs/%v/actions/variables?per_page=100", opts.Owner), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, v := range queryResult.Variables {
				vals = append(vals, v.Name, v.Value) // TODO updated at
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionVariableFields completes variable fields
//
//	name
//	value
func ActionVariableFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"value",
		"visibility",
		"updatedAt",
		"numSelectedRepos",
		"selectedReposURL",
	)
}
