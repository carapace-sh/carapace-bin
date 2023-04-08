package tea

import "github.com/rsteube/carapace"

// ActionIssueFields completes issue fields
//
//	index
//	state
func ActionIssueFields() carapace.Action {
	return carapace.ActionValues(
		"index",
		"state",
		"kind",
		"author",
		"author-id",
		"url",
		"title",
		"body",
		"created",
		"updated",
		"deadline",
		"assignees",
		"milestone",
		"labels",
		"comments",
	)
}
