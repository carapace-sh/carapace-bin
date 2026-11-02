package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type user struct {
	Login string
	Name  string
}

type mentionableUsersQuery struct {
	Data struct {
		Repository struct {
			MentionableUsers struct {
				Nodes []user
			}
		}
	}
}

type assignableUsersQuery struct {
	Data struct {
		Repository struct {
			AssignableUsers struct {
				Nodes []user
			}
		}
	}
}

// ActionMentionableUsers completes mentionable users
//
//	userA (User A)
//	userB (User B)
func ActionMentionableUsers(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult mentionableUsersQuery
		return graphQlAction(opts, fmt.Sprintf(`repository(owner: $owner, name: $repo){ mentionableUsers(first: 100, query: %q) { nodes { login, name } } }`, c.Value), &queryResult, func() carapace.Action {
			users := queryResult.Data.Repository.MentionableUsers.Nodes
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Login
				vals[index*2+1] = user.Name
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("mentionable users")
}

// ActionAssignableUsers completes assignable users
//
//	userA (User A)
//	userB (User B)
func ActionAssignableUsers(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult assignableUsersQuery
		return graphQlAction(opts, fmt.Sprintf(`repository(owner: $owner, name: $repo){ assignableUsers(first: 100, query: %q) { nodes { login, name } } }`, c.Value), &queryResult, func() carapace.Action {
			users := queryResult.Data.Repository.AssignableUsers.Nodes
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Login
				vals[index*2+1] = user.Name
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("assignable users")
}
