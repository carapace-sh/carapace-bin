package tea

import "github.com/rsteube/carapace"

// ActionUserFields completes user fields
//
//	id
//	login
func ActionUserFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"login",
		"full_name",
		"email",
		"avatar_url",
		"language",
		"is_admin",
		"restricted",
		"prohibit_login",
		"location",
		"website",
		"description",
		"visibility",
		"activated",
		"lastlogin_at",
		"created_at",
	)
}
