package winget

import "github.com/carapace-sh/carapace"

// ActionScopes completes scopes
//
//	machine
//	user
func ActionScopes() carapace.Action {
	return carapace.ActionValues(
		"machine",
		"user",
	).Tag("scopes")
}
