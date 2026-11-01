package ghostty

import "github.com/carapace-sh/carapace"

// ActionBellFeatures completes bell features
//
//	audio (Play a custom sound)
//	attention (Request the user's attention when Ghostty is unfocused)
func ActionBellFeatures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"attention", "Request the user's attention when Ghostty is unfocused",
		"audio", "Play a custom sound",
		"border", "Display a border around the alerted surface",
		"system", "Instruct the system to notify the user using built-in system functions",
		"title", "Prepend a bell emoji (🔔) to the title of the alerted surface",
	).Tag("bell features").Uid("ghostty", "bell-feature")
}
