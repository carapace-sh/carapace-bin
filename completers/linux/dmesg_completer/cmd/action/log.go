package action

import "github.com/carapace-sh/carapace"

func ActionLogLevels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"alert", "action must be taken immediately",
		"crit", "critical conditions",
		"debug", "debug-level messages",
		"emerg", "system is unusable",
		"err", "error conditions",
		"info", "informational",
		"notice", "normal but significant condition",
		"warn", "warning conditions",
	)
}
