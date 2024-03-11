package systemctl

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLogLevels completes log levels
//
//	0
//	emerg
func ActionLogLevels() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValuesDescribed(
			"0", "system is unusable", style.Carapace.LogLevelFatal,
			"1", "action must be taken immediately", style.Carapace.LogLevelFatal,
			"2", "critical conditions", style.Carapace.LogLevelCritical,
			"3", "error conditions", style.Carapace.LogLevelError,
			"4", "warning conditions", style.Carapace.LogLevelWarning,
			"5", "normal, but significant, condition", style.Carapace.LogLevelTrace,
			"6", "informational message", style.Carapace.LogLevelInfo,
			"7", "debug-level message", style.Carapace.LogLevelDebug,

			"emerg", "system is unusable", style.Carapace.LogLevelFatal,
			"alert", "action must be taken immediately", style.Carapace.LogLevelFatal,
			"crit", "critical conditions", style.Carapace.LogLevelCritical,
			"err", "error conditions", style.Carapace.LogLevelError,
			"warning", "warning conditions", style.Carapace.LogLevelWarning,
			"notice", "normal, but significant, condition", style.Carapace.LogLevelTrace,
			"info", "informational message", style.Carapace.LogLevelInfo,
			"debug", "debug-level message", style.Carapace.LogLevelDebug,
		).Tag("log levels")
	})
}
