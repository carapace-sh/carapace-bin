package git

import "github.com/carapace-sh/carapace"

// ActionCleanupModes completes cleanup modes
//
//	strip (strip empty lines and trailing whitespace)
//	whitespace (same as strip except #commentary is not removed)
func ActionCleanupModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"default", " Same as strip if the message is to be edited. Otherwise whitespace",
		"scissors", "same as whitespace except that everything from (and including) the line found below is truncated",
		"strip", "strip empty lines and trailing whitespace",
		"verbatim", "do not change the message at all",
		"whitespace", "same as strip except #commentary is not removed",
	).Tag("cleanup modes").Uid("git", "cleanup-mode")
}
