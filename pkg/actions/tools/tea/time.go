package tea

import "github.com/rsteube/carapace"

// ActionTimeFields completes time fields
//
//	id
//	create
func ActionTimeFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"created",
		"repo",
		"issue",
		"user",
		"duration",
	)
}
