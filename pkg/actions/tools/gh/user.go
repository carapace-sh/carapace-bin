package gh

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

type user struct {
	Login string
	Name  string
}

type userQuery struct {
	Data struct {
		Search struct {
			Edges []struct {
				Node user
			}
		}
	}
}

type UserOpts struct {
	Host          string
	Users         bool
	Organizations bool
}

func (o UserOpts) host() string {
	if o.Host == "" {
		return "github.com"
	} else {
		return o.Host
	}
}
func (o UserOpts) owner() string { return "" }
func (o UserOpts) name() string  { return "" }

func (o UserOpts) Default() UserOpts {
	o.Users = true
	o.Organizations = true
	return o
}

func (o UserOpts) format() string {
	filter := make([]string, 0)
	if o.Users {
		filter = append(filter, "... on User { login name }")
	}
	if o.Organizations {
		filter = append(filter, "... on Organization { login name }")
	}
	return strings.Join(filter, " ")
}

// ActionUsers completes users and organizations
//   user (user name)
//   org (organization name)
func ActionUsers(opts UserOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("user search needs at least two characters")
		}

		var queryResult userQuery
		return graphQlAction(opts, fmt.Sprintf(`search(query: "%v in:login", type: USER, first: 100) { edges { node { %v } } }`, c.CallbackValue, opts.format()), &queryResult, func() carapace.Action {
			users := queryResult.Data.Search.Edges
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Node.Login
				vals[index*2+1] = user.Node.Name
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})

}
