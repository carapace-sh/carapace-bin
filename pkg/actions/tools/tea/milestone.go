package tea

import "github.com/carapace-sh/carapace"

// ActionMilestoneFields completes milestone fields
//
//	title
//	state
func ActionMilestoneFields() carapace.Action {
	return carapace.ActionValues(
		"title",
		"state",
		"items_open",
		"items_closed",
		"items",
		"duedate",
		"description",
		"created",
		"updated",
		"closed",
		"id",
	)
}
