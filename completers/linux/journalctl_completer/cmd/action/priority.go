package action

import "github.com/carapace-sh/carapace"

func ActionPriorities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0", "emerg",
		"1", "alert",
		"2", "crit",
		"3", "err",
		"4", "warning",
		"5", "notice",
		"6", "info",
		"7", "debug",
		"alert", "1",
		"crit", "2",
		"debug", "7",
		"emerg", "0",
		"err", "3",
		"info", "6",
		"notice", "5",
		"warning", "4",
	)
}
