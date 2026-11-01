package git

import "github.com/carapace-sh/carapace"

// ActionDirstats completes dirstats
//
//	files (Compute the dirstat numbers by counting the number of files changed)
//	cumulative (Count changes in a child directory for the parent directory as well)
func ActionDirstats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"changes", "Compute the dirstat numbers by counting the lines that have been removed from the source",
		"cumulative", "Count changes in a child directory for the parent directory as well",
		"files", "Compute the dirstat numbers by counting the number of files changed",
		"lines", "Compute the dirstat numbers by doing the regular line-based diff analysis",
	).Uid("git", "dirstat")
}
