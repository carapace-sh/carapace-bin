package git

import "github.com/rsteube/carapace"

// ActionDirstats completes dirstats
//
//	files (Compute the dirstat numbers by counting the number of files changed)
//	cumulative (Count changes in a child directory for the parent directory as well)
func ActionDirstats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"changes", "Compute the dirstat numbers by counting the lines that have been removed from the source",
		"lines", "Compute the dirstat numbers by doing the regular line-based diff analysis",
		"files", "Compute the dirstat numbers by counting the number of files changed",
		"cumulative", "Count changes in a child directory for the parent directory as well",
	)
}
