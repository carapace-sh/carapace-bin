package ghostty

import "github.com/carapace-sh/carapace"

// ActionWindowNewTabPositions completes window tab positions
//
//	"current (Insert the new tab after the currently focused tab)
//	"end (Insert the new tab at the end of the tab list)
func ActionWindowNewTabPositions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"current", "Insert the new tab after the currently focused tab",
		"end", "Insert the new tab at the end of the tab list",
	).Tag("window new tab positions").Uid("ghostty", "window-new-tab-position")
}

// ActionWindowPaddingColors completes window padding colors
//
//	background (The background color specified in `background`)
//	extend (Extend the background color of the nearest grid cell)
func ActionWindowPaddingColors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"background", "The background color specified in `background`",
		"extend", "Extend the background color of the nearest grid cell",
		"extend-always", "Same as \"extend\" but always extends without applying heuristics",
	).Tag("window padding colors").Uid("ghostty", "window-padding-color")
}

// ActionWindowSaveStates completes window save states
//
//	default (will use the default system behavior)
//	never (will never save window state)
func ActionWindowSaveStates() carapace.Action {
	return carapace.ActionValuesDescribed(
		"default", "will use the default system behavior",
		"never", "will never save window state",
		"always", "will always save window state",
	).Tag("window save states").Uid("ghostty", "window-save-state")
}

// ActionWindowThemes completes window themes
//
//	auto (Determine the theme based on the configured terminal background color)
//	system (Use the system theme)
func ActionWindowThemes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auto", "Determine the theme based on the configured terminal background color",
		"system", "Use the system theme",
		"light", "Use the light theme regardless of system theme",
		"dark", "Use the dark theme regardless of system theme",
		"ghostty", "Use the background and foreground colors specified in the Ghostty configuration",
	).Tag("window themes").Uid("ghostty", "window-theme")
}
