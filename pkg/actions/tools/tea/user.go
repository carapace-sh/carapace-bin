package tea

import "github.com/carapace-sh/carapace"

// ActionUserFields completes user fields
//
//	id
//	login
func ActionUserFields() carapace.Action {
	return carapace.ActionValues(
		"activated",
		"avatar_url",
		"created_at",
		"description",
		"email",
		"full_name",
		"id",
		"is_admin",
		"language",
		"lastlogin_at",
		"location",
		"login",
		"prohibit_login",
		"restricted",
		"visibility",
		"website",
	)
}
