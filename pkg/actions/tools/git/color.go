package git

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionColorModes completes color modes
//
//	always
//	auto
func ActionColorModes() carapace.Action {
	return carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword)
}
