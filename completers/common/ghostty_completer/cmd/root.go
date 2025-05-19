package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ghostty"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghostty",
	Short: "a fast, feature-rich, and cross-platform terminal emulator ",
	Long:  "https://ghostty.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("e", "e", "", "run command")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("version", false, "show version")
	addConfigs(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": bridge.ActionCarapaceBin().Split(),
	})
}

func addConfigs(cmd *cobra.Command) {
	cmd.Flags().Bool("_xdg-terminal-exec", false, "Set to true if Ghostty was executed as xdg-terminal-exec on Linux")
	cmd.Flags().String("abnormal-command-exit-runtime", "", "The number of milliseconds of runtime below which we consider a process exit to be abnormal")
	cmd.Flags().String("adjust-box-thickness", "", "Thickness in pixels of box drawing characters")
	cmd.Flags().String("adjust-cell-height", "", "Height of the cell")
	cmd.Flags().String("adjust-cell-width", "", "Witdth of the cell")
	cmd.Flags().String("adjust-cursor-height", "", "Height in pixels of the cursor")
	cmd.Flags().String("adjust-cursor-thickness", "", "Thickness in pixels of the bar cursor and outlined rect cursor")
	cmd.Flags().String("adjust-font-baseline", "", "Distance in pixels from the bottom of the cell to the text baseline")
	cmd.Flags().String("adjust-overline-position", "", "Distance in pixels from the top of the cell to the top of the overline")
	cmd.Flags().String("adjust-overline-thickness", "", "Thickness in pixels of the overline")
	cmd.Flags().String("adjust-strikethrough-position", "", "Distance in pixels from the top of the cell to the top of the strikethrough")
	cmd.Flags().String("adjust-strikethrough-thickness", "", "Thickness in pixels of the strikethrough")
	cmd.Flags().String("adjust-underline-position", "", "Distance in pixels from the top of the cell to the top of the underline")
	cmd.Flags().String("adjust-underline-thickness", "", "Thickness in pixels of the underline")
	cmd.Flags().String("adw-toolbar-style", "", "Determines the appearance of the top and bottom bars when using the Adwaita tab bar")
	cmd.Flags().String("auto-update", "", "Control the auto-update functionality of Ghostty")
	cmd.Flags().String("auto-update-channel", "", "The release channel to use for auto-updates")
	cmd.Flags().String("background-blur-radius", "", "A positive value enables blurring of the background when background-opacity is less than 1")
	cmd.Flags().String("background-opacity", "", "The opacity level (opposite of transparency) of the background")
	cmd.Flags().Bool("bold-is-bright", false, "If `true`, the bold text will use the bright color palette")
	cmd.Flags().String("click-repeat-interval", "", "The time in milliseconds between clicks to consider a click a repeat")
	cmd.Flags().Bool("clipboard-paste-bracketed-safe", false, "If true, bracketed pastes will be considered safe")
	cmd.Flags().Bool("clipboard-paste-protection", false, "Require confirmation before pasting text that appears unsafe")
	cmd.Flags().String("clipboard-read", "", "Whether to allow programs running in the terminal to read from the system clipboard")
	cmd.Flags().Bool("clipboard-trim-trailing-spaces", false, "Trims trailing whitespace on data that is copied to the clipboard")
	cmd.Flags().String("clipboard-write", "", "Whether to allow programs running in the terminal to write to the system clipboard")
	cmd.Flags().Bool("config-default-files", false, "When this is true, the default configuration file paths will be loaded")
	cmd.Flags().StringArray("config-file", nil, "Additional configuration files to read")
	cmd.Flags().Bool("confirm-close-surface", false, "Confirms that a surface should be closed before closing it")
	cmd.Flags().String("copy-on-select", "", "Whether to automatically copy selected text to the clipboard")
	cmd.Flags().Bool("cursor-click-to-move", false, "Enables the ability to move the cursor at prompts")
	cmd.Flags().String("cursor-color", "", "The color of the cursor")
	cmd.Flags().Bool("cursor-invert-fg-bg", false, "Swap the foreground and background colors of the cell under the cursor")
	cmd.Flags().String("cursor-opacity", "", "The opacity level (opposite of transparency) of the cursor")
	cmd.Flags().String("cursor-style", "", "The style of the cursor")
	cmd.Flags().Bool("cursor-style-blink", false, "Sets the default blinking state of the cursor")
	cmd.Flags().String("cursor-text", "", "The color of the text under the cursor")
	cmd.Flags().StringArray("custom-shader", nil, "Custom shaders to run after the default shaders")
	cmd.Flags().String("custom-shader-animation", "", "Controls when custom shaders are animated")
	cmd.Flags().Bool("desktop-notifications", false, "bool")
	cmd.Flags().String("enquiry-response", "", "String to send when we receive `ENQ` (`0x05`) from the command that we are running")
	cmd.Flags().Bool("focus-follows-mouse", false, "If true, when there are multiple split panes, the mouse selects the pane that is focused")
	cmd.Flags().StringArray("font-codepoint-map", nil, "Force one or a range of Unicode codepoints to map to a specific named font")
	cmd.Flags().StringArray("font-family", nil, "The font families to use")
	cmd.Flags().StringArray("font-family-bold", nil, "The font families to use for bold")
	cmd.Flags().StringArray("font-family-bold-italic", nil, "The font families to use for bold italic")
	cmd.Flags().StringArray("font-family-italic", nil, "The font families to use for italic")
	cmd.Flags().StringArray("font-feature", nil, "Apply a font feature")
	cmd.Flags().String("font-size", "", "Font size in points")
	cmd.Flags().String("font-style", "", "The named font style to use")
	cmd.Flags().String("font-style-bold", "", "The named font style to use for bold")
	cmd.Flags().String("font-style-bold-italic", "", "The named font style to use for bold italic")
	cmd.Flags().String("font-style-italic", "", "The named font style to use for italic")
	cmd.Flags().String("font-synthetic-style", "", "Control whether Ghostty should synthesize a style if the requested style is not available")
	cmd.Flags().Bool("font-thicken", false, "Draw fonts with a thicker stroke")
	cmd.Flags().StringArray("font-variation", nil, "Set font variation")
	cmd.Flags().StringArray("font-variation-bold", nil, "Set font variation for bold")
	cmd.Flags().StringArray("font-variation-bold-italic", nil, "Set font variation for bold-italic")
	cmd.Flags().StringArray("font-variation-italic", nil, "Set font variation for italic")
	cmd.Flags().String("freetype-load-flags", "", "FreeType load flags to enable")
	cmd.Flags().String("grapheme-width-method", "", "The method to use for calculating the cell width of a grapheme cluster")
	cmd.Flags().Bool("gtk-adwaita", false, "Adwaita theme support")
	cmd.Flags().String("gtk-single-instance", "", "Run in single-instance mode")
	cmd.Flags().String("gtk-tabs-location", "", "Determines the side of the screen that the GTK tab bar will stick to")
	cmd.Flags().Bool("gtk-titlebar", false, "Display the full GTK titlebar")
	cmd.Flags().Bool("gtk-wide-tabs", false, "Use \"wide\" GTK tabs")
	cmd.Flags().String("image-storage-limit", "", "The total amount of bytes that can be used for image data")
	cmd.Flags().String("initial-command", "", "This is the same as \"command\", but only applies to the first terminal surface created")
	cmd.Flags().Bool("initial-window", false, "Create an initial window when Ghostty is run")
	cmd.Flags().Bool("link-url", false, "Enable URL matching")
	cmd.Flags().String("linux-cgroup", "", "Put every surface (tab, split, window) into a dedicated Linux cgroup")
	cmd.Flags().Bool("linux-cgroup-hard-fail", false, "Let cgroup initialization failure cause exit")
	cmd.Flags().String("linux-cgroup-memory-limit", "", "Memory limit for any individual terminal process")
	cmd.Flags().String("linux-cgroup-processes-limit", "", "Number of processes limit for any individual terminal process")
	cmd.Flags().Bool("macos-auto-secure-input", false, "Automatically enable the \"Secure Input\" feature")
	cmd.Flags().String("macos-icon", "", "Customize the macOS app icon")
	cmd.Flags().String("macos-icon-frame", "", "The material to use for the frame of the macOS app icon")
	cmd.Flags().String("macos-icon-ghost-color", "", "The color of the ghost in the macOS app icon")
	cmd.Flags().String("macos-icon-screen-color", "", "The color of the screen in the macOS app icon")
	cmd.Flags().String("macos-non-native-fullscreen", "", "Non-native fullscreen mode")
	cmd.Flags().String("macos-option-as-alt", "", "Treat option key as alt")
	cmd.Flags().Bool("macos-secure-input-indication", false, "Show a graphical indication when secure input is enabled")
	cmd.Flags().String("macos-titlebar-proxy-icon", "", "Whether the proxy icon in the macOS titlebar is visible")
	cmd.Flags().String("macos-titlebar-style", "", "The style of the macOS titlebar")
	cmd.Flags().Bool("macos-window-shadow", false, "Whether to enable the macOS window shadow")
	cmd.Flags().String("minimum-contrast", "", "The minimum contrast ratio between the foreground and background colors")
	cmd.Flags().Bool("mouse-hide-while-typing", false, "Hide the mouse immediately when typing")
	cmd.Flags().String("mouse-scroll-multiplier", "", "Multiplier for scrolling distance with the mouse wheel")
	cmd.Flags().String("mouse-shift-capture", "", "Determines whether running programs can detect the shift key pressed with a mouse click")
	cmd.Flags().String("osc-color-report-format", "", "Sets the reporting format for OSC sequences that request color information")
	cmd.Flags().String("quick-terminal-animation-duration", "", "Duration (in seconds) of the quick terminal enter and exit animation")
	cmd.Flags().Bool("quick-terminal-autohide", false, "Automatically hide the quick terminal when focus shifts to another window")
	cmd.Flags().String("quick-terminal-position", "", "The position of the \"quick\" terminal window")
	cmd.Flags().String("quick-terminal-screen", "", "The screen where the quick terminal should show up")
	cmd.Flags().Bool("quit-after-last-window-closed", false, "Whether or not to quit after the last surface is closed")
	cmd.Flags().String("quit-after-last-window-closed-delay", "", "Controls how long Ghostty will stay running after the last open surface has been closed")
	cmd.Flags().String("resize-overlay", "", "Controls when resize overlays are shown")
	cmd.Flags().String("resize-overlay-duration", "", "Controls how long the overlay is visible on the screen before it is hidde")
	cmd.Flags().String("resize-overlay-position", "", "Controls the position of the overlay")
	cmd.Flags().String("scrollback-limit", "", "The size of the scrollback buffer in bytes")
	cmd.Flags().String("selection-background", "", "The background color for selection")
	cmd.Flags().String("selection-foreground", "", "The foreground color for selection")
	cmd.Flags().Bool("selection-invert-fg-bg", false, "Swap the foreground and background colors of cells for selection")
	cmd.Flags().String("shell-integration", "", "Whether to enable shell integration auto-injection or not")
	cmd.Flags().String("shell-integration-features", "", "Shell integration features to enable")
	cmd.Flags().String("unfocused-split-fill", "", "The color to dim the unfocused split")
	cmd.Flags().String("unfocused-split-opacity", "", "The opacity level (opposite of transparency) of an unfocused split")
	cmd.Flags().Bool("vt-kam-allowed", false, "Allows the \"KAM\" mode")
	cmd.Flags().Bool("wait-after-command", false, "Keep the terminal open after the command exits")
	cmd.Flags().String("window-colorspace", "", "The colorspace to use for the terminal window")
	cmd.Flags().Bool("window-decoration", false, "Enable window decorations")
	cmd.Flags().String("window-height", "", "The initial window height")
	cmd.Flags().Bool("window-inherit-font-size", false, "Inherit the font size of the previously focused window")
	cmd.Flags().Bool("window-inherit-working-directory", false, "Inherit the working directory of the previously focused window")
	cmd.Flags().String("window-new-tab-position", "", "The position where new tabs are created")
	cmd.Flags().Bool("window-padding-balance", false, "The viewport dimensions are usually not perfectly divisible by the cell size")
	cmd.Flags().String("window-padding-color", "", "The color of the padding area of the window")
	cmd.Flags().String("window-padding-x", "", "Horizontal window padding")
	cmd.Flags().String("window-padding-y", "", "Vertical window padding")
	cmd.Flags().String("window-save-state", "", "Whether to enable saving and restoring window state")
	cmd.Flags().Bool("window-step-resize", false, "Resize the window in discrete increments of the focused surface's cell size")
	cmd.Flags().String("window-theme", "", "The theme to use for the windows")
	cmd.Flags().String("window-title-font-family", "", "The font that will be used for the application's window and tab titles")
	cmd.Flags().Bool("window-vsync", false, "Synchronize rendering with the screen refresh rate")
	cmd.Flags().String("window-width", "", "The initial window width")
	cmd.Flags().String("working-directory", "", "The directory to change to after starting the command")
	cmd.Flags().String("x11-instance-name", "", "This controls the instance name field of the `WM_CLASS` X11 property")

	// TODO use font-families from `+list-fonts`?
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"adw-toolbar-style":       ghostty.ActionAdwToolbarStyles(),
		"auto-update":             ghostty.ActionAutoUpdateModes(),
		"auto-update-channel":     ghostty.ActionReleaseChannels(),
		"clipboard-read":          carapace.ActionValues("ask", "allow", "deny").StyleF(style.ForKeyword),
		"clipboard-write":         carapace.ActionValues("ask", "allow", "deny").StyleF(style.ForKeyword),
		"config-file":             carapace.ActionFiles(),
		"copy-on-select":          ghostty.ActionCopyOnSelectModes(),
		"cursor-color":            color.ActionHexColors(),
		"cursor-style":            ghostty.ActionCursorStyles(),
		"cursor-text":             color.ActionHexColors(),
		"custom-shader":           carapace.ActionFiles(),
		"custom-shader-animation": ghostty.ActionShaderAnimationModes(),
		"font-codepoint-map": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues() // TODO codepoint
			default:
				return os.ActionFontFamilies()
			}
		}),
		"font-family":                os.ActionFontFamilies(),
		"font-family-bold":           os.ActionFontFamilies(),
		"font-family-bold-italic":    os.ActionFontFamilies(),
		"font-family-italic":         os.ActionFontFamilies(),
		"font-feature":               carapace.ActionValues(), // TODO font-features
		"font-style":                 carapace.ActionValues(), // TODO font style
		"font-style-bold":            carapace.ActionValues(), // TODO font style
		"font-style-bold-italic":     carapace.ActionValues(), // TODO font style
		"font-style-italic":          carapace.ActionValues(), // TODO font style
		"font-synthetic-style":       ghostty.ActionFontSyntheticStyles().UniqueList(","),
		"font-variation":             carapace.ActionValues(), // TODO font variation
		"font-variation-bold":        carapace.ActionValues(), // TODO font variation
		"font-variation-bold-italic": carapace.ActionValues(), // TODO font variation
		"font-variation-italic":      carapace.ActionValues(), // TODO font variation
		"freetype-load-flags":        ghostty.ActionFreetypeLoadFlags().UniqueList(","),
		"grapheme-width-method":      ghostty.ActionGraphemeWidthMethods(),
		"gtk-single-instance":        carapace.ActionValues("true", "false", "detect").StyleF(style.ForKeyword),
		"gtk-tabs-location":          carapace.ActionValues("top", "bottom", "left", "right", "hidden"),
		"initial-command":            bridge.ActionCarapaceBin().Split(),
		"linux-cgroup": carapace.ActionValuesDescribed(
			"never", "Never use cgroups",
			"always", "Always use cgroups",
			"single-instance", "Enable cgroups only for Ghostty instances launched",
		).StyleF(style.ForKeyword),
		"macos-icon":                  ghostty.ActionMacIcons(),
		"macos-icon-frame":            ghostty.ActionMacIconFrames(),
		"macos-icon-ghost-color":      color.ActionHexColors(),
		"macos-icon-screen-color":     color.ActionHexColors().List(","),
		"macos-non-native-fullscreen": ghostty.ActionMacFullscreenModes(),
		"macos-option-as-alt":         carapace.ActionValues("false", "true", "left", "right").StyleF(style.ForKeyword),
		"macos-titlebar-proxy-icon":   carapace.ActionValues("visible", "hidden"),
		"macos-titlebar-style":        ghostty.ActionMacTitlebarStyles(),
		"mouse-shift-capture":         ghostty.ActionMouseShiftCaptureModes(),
		"osc-color-report-format":     ghostty.ActionOscColorReportFormats(),
		"quick-terminal-position":     ghostty.ActionQuickTerminalPositions(),
		"quick-terminal-screen":       ghostty.ActionQuickTerminalScreens(),
		"resize-overlay":              ghostty.ActionResizeOverlayModes(),
		"resize-overlay-position":     ghostty.ActionResizeOverlayPositions(),
		"selection-background":        color.ActionHexColors(),
		"selection-foreground":        color.ActionHexColors(),
		"shell-integration":           ghostty.ActionShellIntegrationModes(),
		"shell-integration-features":  ghostty.ActionShellIntegrationFeatures().UniqueList(","),
		"unfocused-split-fill":        color.ActionXtermColorNames(),
		"window-colorspace":           carapace.ActionValues("srgb", "display-p3"),
		"window-new-tab-position":     ghostty.ActionWindowNewTabPositions(),
		"window-padding-color":        ghostty.ActionWindowPaddingColors(),
		"window-save-state":           ghostty.ActionWindowSaveStates(),
		"window-theme":                ghostty.ActionWindowThemes(),
		"window-title-font-family":    os.ActionFontFamilies(),
		"working-directory":           carapace.ActionDirectories(),
	})
}
