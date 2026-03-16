package gh

import "github.com/carapace-sh/carapace"

// ActionSessionFields completes session fields
//
//	id
//	name
func ActionSessionFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"name",
		"state",
		"repository",
		"user",
		"createdAt",
		"updatedAt",
		"completedAt",
		"pullRequestNumber",
		"pullRequestUrl",
		"pullRequestTitle",
		"pullRequestState",
	).Tag("session fields")
}
