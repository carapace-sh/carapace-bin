package gh

import "github.com/carapace-sh/carapace"

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
