package ghostty

import "github.com/carapace-sh/carapace"

// ActionMacIcons completes mac icons
//
//	official (Use the official Ghostty icon)
//	custom-style (Use the official Ghostty icon but with custom styles)
func ActionMacIcons() carapace.Action {
	return carapace.ActionValuesDescribed(
		"official", "Use the official Ghostty icon",
		"blueprint", "Blueprint-style icon with a technical drawing aesthetic",
		"chalkboard", "Chalkboard-style icon with a retro classroom feel",
		"microchip", "Microchip-style icon with a circuit board look",
		"glass", "Glass-style icon with a translucent, modern appearance",
		"holographic", "Holographic-style icon with an iridescent, futuristic look",
		"paper", "Paper-style icon with a folded paper texture",
		"retro", "Retro-style icon with a classic computer aesthetic",
		"xray", "X-ray-style icon showing internal structure",
		"custom", "Use a custom icon from macos-custom-icon path",
		"custom-style", "Use the official Ghostty icon but with custom styles",
	).Tag("mac icons").Uid("ghostty", "macos-icon")
}

// ActionMacIconFrames completes mac icon frames
//
//	aluminum
//	beige
func ActionMacIconFrames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"aluminum", "A brushed aluminum frame. This is the default",
		"beige", "A classic 90's computer beige frame",
		"plastic", "A glossy, dark plastic frame",
		"chrome", "A shiny chrome frame",
	).Tag("mac icon frames").Uid("ghostty", "macos-icon-frame")
}

// ActionMacFullscreenModes completes mac fullscreen modes
//
//	visible-menu (Use non-native macOS fullscreen, keep the menu bar visible)
//	true (Use non-native macOS fullscreen, hide the menu bar)
func ActionMacFullscreenModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"visible-menu", "Use non-native macOS fullscreen, keep the menu bar visible",
		"true", "Use non-native macOS fullscreen, hide the menu bar",
		"false", "Use native macOS fullscreen",
		"padded-notch", "Non-native fullscreen with padding for the notch area",
	).Tag("mac fullscreen modes").Uid("ghostty", "macos-fullscreen-mode")
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
	).Tag("mac titlebar styles").Uid("ghostty", "macos-titlebar-style")
}
