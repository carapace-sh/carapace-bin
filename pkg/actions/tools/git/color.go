package git

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionColorModes completes color modes
//
//	always
//	auto
func ActionColorModes() carapace.Action {
	return carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword)
}
