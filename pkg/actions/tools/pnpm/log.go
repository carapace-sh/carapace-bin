package pnpm

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLoglevels completes log levels
//
//	debug
//	info
func ActionLoglevels() carapace.Action {
	return carapace.ActionValues("debug", "info", "warn", "error").StyleF(style.ForLogLevel)
}
