package git

import "github.com/carapace-sh/carapace"

// ActionMergeStrategies completes merge strategies
//
//	octopus (resolve cases with more than two heads)
//	ours (auto-resolve cleanly by favoring our version)
func ActionMergeStrategies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"octopus", "resolve cases with more than two heads",
		"ort", "replace recursive with a new algorithm",
		"ours", "auto-resolve cleanly by favoring our version",
		"recursive", "alias for ort (deprecated)",
		"resolve", "resolve two heads using a 3-way merge algorithm",
		"subtree", "modified recursive straty with tree adjustment",
	).Tag("merge strategies").Uid("git", "merge-strategy")
}

// ActionMergeStrategyOptions completes merge strategy options
//
//	ours (auto-resolve favoring ours)
//	theirs (auto-resolve favoring theirs)
func ActionMergeStrategyOptions(strategy string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch strategy {
		case "ort", "recursive":
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.Batch(
						carapace.ActionValuesDescribed(
							"find-renames", "turn on rename detection",
							"ignore-all-space", "ignore all space",
							"ignore-cr-at-eol", "ignore <cr> at end of line",
							"ignore-space-at-eol", "ignore <space> at end of line",
							"ignore-space-change", "ignore space change",
							"no-renames", "turn off rename detection",
							"no-renormalize", "disable renormalize",
							"ours", "auto-resolve favoring ours",
							"patience", "spend extra time to avoid mismerges",
							"renormalize", "enable renormalize",
							"subtree", "advance subtree stratebgy",
							"theirs", "auto-resolve favoring theirs",
						).Uid("git", "merge-strategy-option"),
						carapace.ActionValuesDescribed(
							"diff-algorithm", "set diff allgorithm",
						).Suffix("=").Uid("git", "merge-strategy-option"),
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
