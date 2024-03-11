package ffmpeg

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLogLevels completes log levels
//
//	verbose (Same as "info", except more verbose)
//	warning (Show all warnings and errors)
func ActionLogLevels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"quiet", "Show nothing at all; be silent",
		"panic", "Only show fatal errors which could lead the process to crash",
		"fatal", "Only show fatal errors",
		"error", "Show all errors, including ones which can be recovered from",
		"warning", "Show all warnings and errors",
		"info", "Show informative messages during processing",
		"verbose", "Same as \"info\", except more verbose",
		"debug", "Show everything, including debugging information",
		"trace", "",
	).StyleF(style.ForLogLevel)
}
