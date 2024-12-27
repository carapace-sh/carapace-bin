package ghostty

import "github.com/carapace-sh/carapace"

// ActionMacIcons completes mac icons
func ActionMacIcons() carapace.Action {
	return carapace.ActionValuesDescribed(
		"official", "Use the official Ghostty icon",
		"custom-style", "Use the official Ghostty icon but with custom styles",
	).Tag("mac icons")
}

// ActionMacIconFrames completes mac icon frames
func ActionMacIconFrames() carapace.Action {
	return carapace.ActionValues(
		"aluminum", "A brushed aluminum frame. This is the default",
		"beige", "A classic 90's computer beige frame",
		"plastic", "A glossy, dark plastic frame",
		"chrome", "A shiny chrome frame",
	).Tag("mac icon frames")
}

// ActionMacFullscreenModes completes mac fullscreen modes
func ActionMacFullscreenModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"visible-menu", "Use non-native macOS fullscreen, keep the menu bar visible",
		"true", "Use non-native macOS fullscreen, hide the menu bar",
		"false", "Use native macOS fullscreen",
	).Tag("mac fullscreen modes")
}

// ActionMacTitlebarStyles completes mac titlebar styles
//
//	native (native macOS titlebar with zero customization)
//	transparent (same as "native" but the titlebar will be transparent)
func ActionMacTitlebarStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"native", "native macOS titlebar with zero customization",
		"transparent", "same as \"native\" but the titlebar will be transparent",
		"tabs", "custom titlebar that integrates the tab bar into the titlebar",
		"hidden", "hides the titlebar",
	).Tag("mac titlebar styles")
}
