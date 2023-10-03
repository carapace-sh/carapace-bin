package gh

import "github.com/rsteube/carapace"

// ActionDeployKeyFields completes deploy key fields
//
//	id
//	key
func ActionDeployKeyFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"key",
		"title",
		"createdAt",
		"readOnly",
	)
}
