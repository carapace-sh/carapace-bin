package terramate

import "github.com/rsteube/carapace"

// ActionSafeguards completes safeguards
//
//	all (Disable all checks)
//	git (Disable all git related checks)
func ActionSafeguards() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "Disable all checks",
		"git", "Disable all git related checks",
		"git-out-of-sync", "Disable the check for git remote out of sync",
		"git-uncommitted", "Disable the check for uncommitted files",
		"git-untracked", "Disable the check for untracked files",
		"outdated-code", "Disable the check for outdated code",
	)
}
