package tea

import "github.com/carapace-sh/carapace"

// ActionMilestoneFields completes milestone fields
//
//	title
//	state
func ActionMilestoneFields() carapace.Action {
	return carapace.ActionValues(
		"closed",
		"created",
		"description",
		"duedate",
		"id",
		"items",
		"items_closed",
		"items_open",
		"state",
		"title",
		"updated",
	)
}
