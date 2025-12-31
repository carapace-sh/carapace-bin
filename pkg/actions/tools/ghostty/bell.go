package ghostty

import "github.com/carapace-sh/carapace"

// ActionBellFeatures completes bell features
//
//	audio (Play a custom sound)
//	attention (Request the user's attention when Ghostty is unfocused)
func ActionBellFeatures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"system", "Instruct the system to notify the user using built-in system functions",
		"audio", "Play a custom sound",
		"attention", "Request the user's attention when Ghostty is unfocused",
		"title", "Prepend a bell emoji (🔔) to the title of the alerted surface",
		"border", "Display a border around the alerted surface",
	).Tag("bell features")
}
