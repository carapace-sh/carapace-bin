package git

import "github.com/carapace-sh/carapace"

// ActionMergeStrategies completes merge strategies
//
//	octopus (resolve cases with more than two heads)
//	ours (auto-resolve cleanly by favoring our version)
func ActionMergeStrategies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"octopus", "resolve cases with more than two heads",
		"ours", "auto-resolve cleanly by favoring our version",
		"recursive", "recursively resolve two heads using a 3-way merge algorithm",
		"resolve", "resolve two heads using a 3-way merge algorithm",
		"subtree", "modified recursive straty with tree adjustment",
	).Tag("merge strategies")
}

// ActionMergeStrategyOptions completes merge strategy options
//
//	ours (auto-resolve favoring ours)
//	theirs (auto-resolve favoring theirs)
func ActionMergeStrategyOptions(strategy string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch strategy {
		case "recursive":
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.Batch(
						carapace.ActionValuesDescribed(
							"ours", "auto-resolve favoring ours",
							"theirs", "auto-resolve favoring theirs",
							"patience", "spend extra time to avoid mismerges",
							"ignore-space-change", "ignore space change",
							"ignore-all-space", "ignore all space",
							"ignore-space-at-eol", "ignore <space> at end of line",
							"ignore-cr-at-eol", "ignore <cr> at end of line",
							"renormalize", "enable renormalize",
							"no-renormalize", "disable renormalize",
							"no-renames", "turn off rename detection",
							"find-renames", "turn on rename detection",
							"subtree", "advance subtree stratebgy",
						),
						carapace.ActionValuesDescribed(
							"diff-algorithm", "set diff allgorithm",
						).Suffix("="),
					).ToA()
				default:
					switch c.Parts[0] {
					case "diff-algorithm":
						return ActionDiffAlgorithms()
					default:
						return carapace.ActionValues()
					}
				}
			})
		default:
			return carapace.ActionValues()
		}
	})
}
