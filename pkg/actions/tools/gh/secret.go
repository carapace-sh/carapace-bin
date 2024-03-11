package gh

import "github.com/carapace-sh/carapace"

// ActionSecretFields completes secret fields
//
//	name
//	visibility
func ActionSecretFields() carapace.Action {
	return carapace.ActionValues(
		"selectedReposURL",
		"name",
		"visibility",
		"updatedAt",
		"numSelectedRepos",
	)
}
