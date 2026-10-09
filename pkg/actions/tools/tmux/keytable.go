package tmux

import (
	"github.com/carapace-sh/carapace"
)

// ActionKeyTables completes key tables
//
//	root
//	prefix
func ActionKeyTables() carapace.Action {
	return carapace.ActionValues(
		"prefix",
		"root",
	)
}
