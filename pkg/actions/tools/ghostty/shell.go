package ghostty

import "github.com/carapace-sh/carapace"

// ActionShellIntegrationModes completes shell integration modes
//
//	none (Do not do any automatic injection)
//	detect (Detect the shell based on the filename)
func ActionShellIntegrationModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "Do not do any automatic injection",
		"detect", "Detect the shell based on the filename",
		"bash", "",
		"elvish", "",
		"fish", "",
		"zsh", "",
	).Tag("shell integration modes")
}

// ActionShellIntegrationFeatures completes shell integration features
//
//	cursor (Set the cursor to a blinking bar at the prompt)
//	sudo (Set sudo wrapper to preserve terminfo)
func ActionShellIntegrationFeatures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cursor", "Set the cursor to a blinking bar at the prompt",
		"sudo", "Set sudo wrapper to preserve terminfo",
		"title", "Set the window title via shell integration",
	).Tag("shell integration features")
}
