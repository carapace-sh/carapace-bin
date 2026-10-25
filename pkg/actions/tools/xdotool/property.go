package xdotool

import "github.com/carapace-sh/carapace"

// ActionProperties completes properties
//
//	ABOVE (Show window above all others (always on top))
//	BELOW (Show window below all others)
func ActionProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ABOVE", "Show window above all others (always on top)",
		"BELOW", "Show window below all others",
		"DEMANDS_ATTENTION", "marks window urgent or needing attention",
		"FULLSCREEN", "makes window fullscreen",
		"HIDDEN", "unmaps the window",
		"MAXIMIZED_HORZ", "sizes the window maximized horizontally",
		"MAXIMIZED_VERT", "sizes the window maximized vertically",
		"MODAL", "makes the window into a modal",
		"SHADED", "rolls the window up",
		"SKIP_PAGER", "hides the window from the window pager",
		"SKIP_TASKBAR", "hides the window from the taskbar",
		"STICKY", "makes the window appear on all workspaces",
	)
}
