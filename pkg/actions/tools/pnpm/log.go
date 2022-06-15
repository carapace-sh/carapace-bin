package pnpm

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionLoglevel completes log levels
//
//	debug
//	info
func ActionLoglevel() carapace.Action {
	return carapace.ActionValues("debug", "info", "warn", "error").StyleF(style.ForLogLevel)
}
