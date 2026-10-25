package gh

import "github.com/carapace-sh/carapace"

// ActionSessionFields completes session fields
//
//	id
//	name
func ActionSessionFields() carapace.Action {
	return carapace.ActionValues(
		"completedAt",
		"createdAt",
		"id",
		"name",
		"pullRequestNumber",
		"pullRequestState",
		"pullRequestTitle",
		"pullRequestUrl",
		"repository",
		"state",
		"updatedAt",
		"user",
	).Tag("session fields")
}
