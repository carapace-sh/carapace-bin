package git

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionColumnLayoutModes completions column layout modes
//
//	always (always show in columns)
//	auto (show in columns if the output is to the terminal)
func ActionColumnLayoutModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "always show in columns",
		"auto", "show in columns if the output is to the terminal",
		"column", "fill columns before rows",
		"dense", "make unequal size columns to utilize more space",
		"never", "never show in columns",
		"nodense", "make equal size columns",
		"plain", "show in one column",
		"row", "fill rows before columns",
	).StyleF(style.ForKeyword)
}
