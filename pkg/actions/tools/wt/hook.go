package wt

import "github.com/carapace-sh/carapace"

// ActionHookTypes completes hook types
//
//	post-commit
//	post-create
func ActionHookTypes() carapace.Action {
	return carapace.ActionValues(
		"post-commit",
		"post-create",
		"post-merge",
		"post-remove",
		"post-start",
		"post-switch",
		"pre-commit",
		"pre-merge",
		"pre-remove",
		"pre-switch",
	).Tag("hook types")
}
