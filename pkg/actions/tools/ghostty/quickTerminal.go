package ghostty

import "github.com/carapace-sh/carapace"

// ActionQuickTerminalPositions completes quick terminal positions
//
//	top (Terminal appears at the top of the screen)
//	bottom (Terminal appears at the bottom of the screen)
func ActionQuickTerminalPositions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"top", "Terminal appears at the top of the screen",
		"bottom", "Terminal appears at the bottom of the screen",
		"left", "Terminal appears at the left of the screen",
		"right", "Terminal appears at the right of the screen",
		"center", "Terminal appears at the center of the screen",
	).Tag("quick terminal positions")
}

// ActionQuickTerminalScreens completes quick terminal screens
//
//	main (The screen that the operating system recommends as the main screen)
//	mouse (The screen that the mouse is currently hovered over)
func ActionQuickTerminalScreens() carapace.Action {
	return carapace.ActionValuesDescribed(
		"main", "The screen that the operating system recommends as the main screen",
		"mouse", "The screen that the mouse is currently hovered over",
		"macos-menu-bar", "The screen that contains the macOS menu bar",
	).Tag("quick terminal screens")
}
