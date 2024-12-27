package ghostty

import "github.com/carapace-sh/carapace"

// ActionAutoUpdateModes completes auto update modes
//
//	off (Disable auto-updates)
//	check (Check for updates)
func ActionAutoUpdateModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"off", "Disable auto-updates",
		"check", "Check for updates",
		"download", "Check for updates, automatically download the update",
	).Tag("auto-update modes")
}

// ActionReleaseChannels completes release channels
//
//	stable (Stable, tagged releases such as \"1.0.0\")
//	tip (Pre-release versions generated from each commit to the main branch)
func ActionReleaseChannels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"stable", "Stable, tagged releases such as \"1.0.0\"",
		"tip", "Pre-release versions generated from each commit to the main branch",
	).Tag("release channels")
}
