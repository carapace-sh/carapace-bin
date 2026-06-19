package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionNotifyOnCommandFinish completes notify on command finish modes
//
//	never (Never send notifications)
//	unfocused (Only send notifications if the surface is not focused)
func ActionNotifyOnCommandFinish() carapace.Action {
	return carapace.ActionValuesDescribed(
		"never", "Never send notifications",
		"unfocused", "Only send notifications if the surface is not focused",
		"always", "Always send notifications",
	).StyleF(style.ForKeyword).Tag("notify on command finish modes")
}

// ActionNotifyOnCommandFinishActions completes notify on command finish actions
//
//	bell (enabled by default)
//	notify (disabled by default)
func ActionNotifyOnCommandFinishActions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bell", "enabled by default",
		"notify", "disabled by default",
		"no-bell", "disabled by default",
		"no-notify", "disabled by default",
	).Tag("notify on command finish actions")
}

// ActionRightClickActions completes right click actions
//
//	context-menu (Show the context menu)
//	paste (Paste the contents of the clipboard)
func ActionRightClickActions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"context-menu", "Show the context menu",
		"paste", "Paste the contents of the clipboard",
		"copy", "Copy the selected text to the clipboard",
		"copy-or-paste", "If there is a selection copy; otherwise paste",
		"ignore", "Do nothing, ignore the right-click",
	).Tag("right click actions")
}

// ActionScrollbarModes completes scrollbar modes
//
//	system (Respect the system settings)
//	never (Never show a scrollbar)
func ActionScrollbarModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"system", "Respect the system settings for when to show scrollbars",
		"never", "Never show a scrollbar",
	).StyleF(style.ForKeyword).Tag("scrollbar modes")
}

// ActionWindowShowTabBar completes window show tab bar modes
//
//	always (Always display the tab bar)
//	auto (Automatically show and hide the tab bar)
func ActionWindowShowTabBar() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "Always display the tab bar, even with only one tab",
		"auto", "Automatically show and hide the tab bar",
		"never", "Never show the tab bar",
	).StyleF(style.ForKeyword).Tag("window show tab bar modes")
}

// ActionMacosDockDropBehaviors completes macos dock drop behaviors
//
//	new-tab (Create a new tab in the current window)
//	new-window (Create a new window unconditionally)
func ActionMacosDockDropBehaviors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"new-tab", "Create a new tab in the current window, or open a new window if none exist",
		"new-window", "Create a new window unconditionally",
	).Tag("macos dock drop behaviors")
}

// ActionMacosHidden completes macos hidden modes
//
//	never (The macOS app is never hidden)
//	always (The macOS app is always hidden)
func ActionMacosHidden() carapace.Action {
	return carapace.ActionValuesDescribed(
		"never", "The macOS app is never hidden",
		"always", "The macOS app is always hidden",
	).Tag("macos hidden modes")
}

// ActionMacosShortcuts completes macos shortcuts modes
//
//	ask (Ask the user for permission)
//	allow (Allow Shortcuts to control Ghostty without asking)
func ActionMacosShortcuts() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ask", "Ask the user for permission",
		"allow", "Allow Shortcuts to control Ghostty without asking",
		"deny", "Deny Shortcuts from controlling Ghostty",
	).StyleF(style.ForKeyword).Tag("macos shortcuts modes")
}

// ActionMacosWindowButtons completes macos window buttons visibility
//
//	visible (Show the window buttons)
//	hidden (Hide the window buttons)
func ActionMacosWindowButtons() carapace.Action {
	return carapace.ActionValuesDescribed(
		"visible", "Show the window buttons",
		"hidden", "Hide the window buttons",
	).Tag("macos window buttons visibility")
}

// ActionGtkTitlebarStyles completes gtk titlebar styles
//
//	native (traditional titlebar with a title)
//	tabs (merges the tab bar and the traditional titlebar)
func ActionGtkTitlebarStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"native", "traditional titlebar with a title, a few buttons and window controls",
		"tabs", "merges the tab bar and the traditional titlebar",
	).Tag("gtk titlebar styles")
}

// ActionGtkQuickTerminalLayers completes gtk quick terminal layers
//
//	overlay (appears in front of all windows)
//	top (appears in front of normal windows but behind fullscreen overlays)
func ActionGtkQuickTerminalLayers() carapace.Action {
	return carapace.ActionValuesDescribed(
		"overlay", "The quick terminal appears in front of all windows",
		"top", "The quick terminal appears in front of normal windows but behind fullscreen overlays",
		"bottom", "The quick terminal appears behind normal windows but in front of wallpapers",
		"background", "The quick terminal appears behind all windows",
	).Tag("gtk quick terminal layers")
}

// ActionQuickTerminalSpaceBehaviors completes quick terminal space behaviors
//
//	move (quick terminal moved to the current space)
//	remain (quick terminal stays in the original space)
func ActionQuickTerminalSpaceBehaviors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"move", "The quick terminal will move to the current space when switching",
		"remain", "The quick terminal will stay in the space where it was originally opened",
	).Tag("quick terminal space behaviors")
}

// ActionSplitPreserveZoom completes split preserve zoom modes
//
//	navigation (preserve the zoomed split state when navigating)
func ActionSplitPreserveZoom() carapace.Action {
	return carapace.ActionValuesDescribed(
		"navigation", "preserve the zoomed split state when navigating to another split",
		"no-navigation", "unzoom when navigating to another split",
	).Tag("split preserve zoom modes")
}

// ActionWindowSubtitle completes window subtitle modes
//
//	false (Disable the subtitle)
//	working-directory (Set the subtitle to the working directory)
func ActionWindowSubtitle() carapace.Action {
	return carapace.ActionValuesDescribed(
		"false", "Disable the subtitle",
		"working-directory", "Set the subtitle to the working directory of the surface",
	).Tag("window subtitle modes")
}
