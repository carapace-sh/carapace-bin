package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type emailQuery struct {
	Data struct {
		User struct {
			Name  string
			Email string
		}
		Organization struct {
			Name  string
			Email string
		}
	}
}

// ActionUserEmails completes email for given user
func ActionUserEmails(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult emailQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`user(login: %#v) { name email }`, opts.Owner), &queryResult, func() carapace.Action {
			if queryResult.Data.User.Email == "" {
				queryResult.Data.User.Email = fmt.Sprintf(`%v@users.noreply.github.com`, opts.Owner)
			}
			return carapace.ActionValuesDescribed(queryResult.Data.User.Email, queryResult.Data.User.Email)
		})
	})
}

// ActionOrganizationEmails completes email for given organization
func ActionOrganizationEmails(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult emailQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`organization(login: %#v) { name email }`, opts.Owner), &queryResult, func() carapace.Action {
			if queryResult.Data.Organization.Email == "" {
				queryResult.Data.Organization.Email = fmt.Sprintf(`%v@organizations.noreply.github.com`, opts.Owner)
			}
			return carapace.ActionValuesDescribed(queryResult.Data.Organization.Email, queryResult.Data.Organization.Email)
		})
	})
}
