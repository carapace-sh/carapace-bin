package doing

import "github.com/carapace-sh/carapace"

// ActionWorkItemTypes completes work item types
//
//	Bug
//	Task
func ActionWorkItemTypes() carapace.Action {
	return carapace.ActionValues(
		"Bug",
		"Epic",
		"Feature",
		"Issue",
		"Task",
		"Test Case",
		"User Story",
	)
}
