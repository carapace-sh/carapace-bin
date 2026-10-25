package gh

import "github.com/carapace-sh/carapace"

// ActionDeployKeyFields completes deploy key fields
//
//	id
//	key
func ActionDeployKeyFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"id",
		"key",
		"readOnly",
		"title",
	)
}
