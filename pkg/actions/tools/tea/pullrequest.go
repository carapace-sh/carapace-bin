package tea

import "github.com/rsteube/carapace"

// ActionPullrequestFields completes pullrequest fields
//
//	index
//	state
func ActionPullrequestFields() carapace.Action {
	return carapace.ActionValues(
		"index",
		"state",
		"author",
		"author-id",
		"url",
		"title",
		"body",
		"mergeable",
		"base",
		"base-commit",
		"head",
		"diff",
		"patch",
		"created",
		"updated",
		"deadline",
		"assignees",
		"milestone",
		"labels",
		"comments",
	)
}
