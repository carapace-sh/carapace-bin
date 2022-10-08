package kubectl

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionValidationModes completes validation modes
//
//	strict
//	warn
func ActionValidationModes() carapace.Action {
	return carapace.ActionValues("strict", "true", "warn", "ignore", "false").StyleF(style.ForKeyword)
}
