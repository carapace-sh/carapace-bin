package golangcilint

import "github.com/rsteube/carapace"

// ActionPresets completes presets
//
//	comment
//	complexity
func ActionPresets() carapace.Action {
	return carapace.ActionValues("bugs",
		"comment",
		"complexity",
		"error",
		"format",
		"import",
		"metalinter",
		"module",
		"performance",
		"sql",
		"style",
		"test",
		"unused",
	)
}
