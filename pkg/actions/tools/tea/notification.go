package tea

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionNotificationFields completes notification fields
//
//	id
//	status
func ActionNotificationFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"status",
		"updated",
		"index",
		"type",
		"state",
		"title",
		"repository",
	)
}

// ActionNotificationStates complets notification states
//
//	pinned
//	unread
func ActionNotificationStates() carapace.Action {
	return carapace.ActionStyledValues(
		"pinned", style.Carapace.KeywordAmbiguous,
		"unread", style.Carapace.KeywordUnknown,
		"read", style.Carapace.KeywordPositive,
	)
}

// ActionNotificationTypes completes notification types
//
//	issue
//	pull
func ActionNotificationTypes() carapace.Action {
	return carapace.ActionValues(
		"issue",
		"pull",
		"repository",
		"commit",
	)

}
