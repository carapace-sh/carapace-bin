package git

import "github.com/carapace-sh/carapace"

// ActionUpdateRefsModes completes update-refs modes
//
//	branches (all local branches will be rewritten)
//	head (only the current HEAD reference will be rewritten)
func ActionUpdateRefsModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"branches", "all local branches will be rewritten",
		"head", "only the current HEAD reference will be rewritten",
	).Uid("git", "update-refs-mode")
}
