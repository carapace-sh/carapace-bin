package gh

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

type owner struct {
	Login string
	Name  string
}

type ownerQuery struct {
	Data struct {
		Search struct {
			Edges []struct {
				Node owner
			}
		}
	}
}

// ActionOwners completes users and organizations
//   user (user name)
//   org (organization name)
func ActionOwners(opts HostOpts) carapace.Action {
	return carapace.Batch(
		ActionUsers(opts),
		ActionOrganizations(opts),
	).ToA()
}

// ActionUsers completes users
//   user (user name)
//   another (another name)
func ActionUsers(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("user search needs at least two characters")
		}

		var queryResult ownerQuery
		return graphQlAction(opts, fmt.Sprintf(`search(query: "%v in:login", type: USER, first: 100) { edges { node { ... on User { login name } } } }`, c.CallbackValue), &queryResult, func() carapace.Action {
			users := queryResult.Data.Search.Edges
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Node.Login
				vals[index*2+1] = user.Node.Name
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Gh.OwnerUser)
		})
	})

}

// ActionOrganizations completes organizations
//   org (organization name)
//   another (another name)
func ActionOrganizations(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("user search needs at least two characters")
		}

		var queryResult ownerQuery
		return graphQlAction(opts, fmt.Sprintf(`search(query: "%v in:login", type: USER, first: 100) { edges { node { ... on Organization { login name } } } }`, c.CallbackValue), &queryResult, func() carapace.Action {
			users := queryResult.Data.Search.Edges
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Node.Login
				vals[index*2+1] = user.Node.Name
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Gh.OwnerOrganization)
		})
	})

}
