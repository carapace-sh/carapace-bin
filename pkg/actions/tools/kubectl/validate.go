package kubectl

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionValidationModes completes validation modes
//
//	strict
//	warn
func ActionValidationModes() carapace.Action {
	return carapace.ActionValues("false", "ignore", "strict", "true", "warn").StyleF(style.ForKeyword)
}
