package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLinuxCgroups completes linux cgroup modes
//
//	never (Never use cgroups)
//	always (Always use cgroups)
func ActionLinuxCgroups() carapace.Action {
	return carapace.ActionValuesDescribed(
		"never", "Never use cgroups",
		"always", "Always use cgroups",
		"single-instance", "Enable cgroups only for Ghostty instances launched",
	).StyleF(style.ForKeyword).Tag("linux cgroup modes").Uid("ghostty", "linux-cgroup")
}

// ActionBackgroundBlurs completes background blur modes
//
//	false (equivalent to a blur intensity of 0)
//	true (equivalent to the default blur intensity of 20)
func ActionBackgroundBlurs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"false", "equivalent to a blur intensity of 0",
		"true", "equivalent to the default blur intensity of 20",
	).StyleF(style.ForKeyword).Tag("background blur modes").Uid("ghostty", "background-blur")
}

// ActionFontShapingBreaks completes font shaping break modes
//
//	cursor (Break runs under the cursor)
func ActionFontShapingBreaks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cursor", "Break runs under the cursor",
	).Tag("font shaping breaks").Uid("ghostty", "font-shaping-break")
}

// ActionGtkSingleInstances completes gtk single instance modes
//
//	true (Run in single-instance mode)
//	false (Do not run in single-instance mode)
func ActionGtkSingleInstances() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "Run in single-instance mode",
		"false", "Do not run in single-instance mode",
		"detect", "Automatically detect whether to use single-instance mode",
	).StyleF(style.ForKeyword).Tag("gtk single instance modes").Uid("ghostty", "gtk-single-instance")
}

// ActionGtkTabsLocations completes gtk tabs locations
//
//	top
//	bottom
func ActionGtkTabsLocations() carapace.Action {
	return carapace.ActionValues("top", "bottom", "left", "right", "hidden").Tag("gtk tabs locations").Uid("ghostty", "gtk-tabs-location")
}

// ActionMacosOptionsAsAlt completes macos option-as-alt modes
//
//	false (Do not treat the Option key as Alt)
//	true (Treat both left and right Option keys as Alt)
func ActionMacosOptionsAsAlt() carapace.Action {
	return carapace.ActionValuesDescribed(
		"false", "Do not treat the Option key as Alt",
		"true", "Treat both left and right Option keys as Alt",
		"left", "Treat only the left Option key as Alt",
		"right", "Treat only the right Option key as Alt",
	).StyleF(style.ForKeyword).Tag("macos option-as-alt modes").Uid("ghostty", "macos-option-as-alt")
}

// ActionMacosTitlebarProxyIcons completes macos titlebar proxy icon visibility
//
//	visible (Show the proxy icon)
//	hidden (Hide the proxy icon)
func ActionMacosTitlebarProxyIcons() carapace.Action {
	return carapace.ActionValuesDescribed(
		"visible", "Show the proxy icon in the macOS titlebar",
		"hidden", "Hide the proxy icon in the macOS titlebar",
	).Tag("macos titlebar proxy icon visibility").Uid("ghostty", "macos-titlebar-proxy-icon")
}

// ActionQuickTerminalKeyboardInteractivities completes quick terminal keyboard interactivity modes
//
//	none (The quick terminal will not receive any keyboard input)
//	on-demand (The quick terminal would only receive keyboard input when it is focused)
func ActionQuickTerminalKeyboardInteractivities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "The quick terminal will not receive any keyboard input",
		"on-demand", "The quick terminal would only receive keyboard input when it is focused",
		"exclusive", "The quick terminal will always receive keyboard input",
	).StyleF(style.ForKeyword).Tag("quick terminal keyboard interactivities").Uid("ghostty", "quick-terminal-keyboard-interactivity")
}

// ActionScrollToBottoms completes scroll-to-bottom modes
//
//	keystroke (scroll the surface to the bottom when the user presses a key)
//	output (scroll the surface to the bottom if there is new data to display)
func ActionScrollToBottoms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"keystroke", "scroll the surface to the bottom when the user presses a key",
		"output", "scroll the surface to the bottom if there is new data to display",
	).Tag("scroll-to-bottom modes").Uid("ghostty", "scroll-to-bottom")
}

// ActionWindowColorspaces completes window colorspaces
//
//	srgb
//	display-p3
func ActionWindowColorspaces() carapace.Action {
	return carapace.ActionValues("srgb", "display-p3").Tag("window colorspaces").Uid("ghostty", "window-colorspace")
}

// ActionTerms completes TERM values
//
//	xterm-ghostty
//	xterm-256color
func ActionTerms() carapace.Action {
	return carapace.ActionValues("xterm-ghostty", "xterm-256color", "xterm-kitty", "xterm-terminfo").Tag("term values").Uid("ghostty", "term")
}
