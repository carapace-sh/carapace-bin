package bridge

import (
	"github.com/rsteube/carapace"
)

// ActionPosener (deprecated and renamed to ActionComplete)
// TODO remove action
func ActionPosener(cmd string) carapace.Action {
	return carapace.ActionMessage("macro renamed to bridge.Complete")
}
