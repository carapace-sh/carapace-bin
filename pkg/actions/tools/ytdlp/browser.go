package ytdlp

import "github.com/rsteube/carapace"

// ActionBrowsers completes browsers
//
//	brave
//	chrome
func ActionBrowsers() carapace.Action {
	return carapace.ActionValues(
		"brave",
		"chrome",
		"chromium",
		"edge",
		"firefox",
		"opera",
		"safari",
		"vivaldi",
	).Tag("browsers")
}
