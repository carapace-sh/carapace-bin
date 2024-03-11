package tea

import "github.com/carapace-sh/carapace"

// ActionRepositoryFields completes repository fields
//
//	description
//	forks
func ActionRepositoryFields() carapace.Action {
	return carapace.ActionValues(
		"description",
		"forks",
		"id",
		"name",
		"owner",
		"stars",
		"ssh",
		"updated",
		"url",
		"permission",
		"type",
	)
}
