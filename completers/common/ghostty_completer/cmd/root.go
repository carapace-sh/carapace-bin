package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ghostty"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("e", "e", false, "run command") // if set positional arguments accept a command
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("version", false, "show version")
	addConfigs(rootCmd)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("e").Changed {
				return bridge.ActionCarapaceBin()
			}
			return carapace.ActionValues()
		}),
	)
}

func addConfigs(cmd *cobra.Command) {
	cmd.Flags().String("abnormal-command-exit-runtime", "", "The number of milliseconds of runtime below which we consider a process exit to be abnormal")
	cmd.Flags().String("adjust-box-thickness", "", "Thickness in pixels of box drawing characters")
	cmd.Flags().String("adjust-cell-height", "", "Height of the cell")
	cmd.Flags().String("adjust-cell-width", "", "Witdth of the cell")
	cmd.Flags().String("adjust-cursor-height", "", "Height in pixels of the cursor")
	cmd.Flags().String("adjust-cursor-thickness", "", "Thickness in pixels of the bar cursor and outlined rect cursor")
	cmd.Flags().String("adjust-font-baseline", "", "Distance in pixels from the bottom of the cell to the text baseline")
	cmd.Flags().String("adjust-icon-height", "", "Height in pixels or percentage adjustment of maximum height for nerd font icons")
	cmd.Flags().String("adjust-overline-position", "", "Distance in pixels from the top of the cell to the top of the overline")
	cmd.Flags().String("adjust-overline-thickness", "", "Thickness in pixels of the overline")
	cmd.Flags().String("adjust-strikethrough-position", "", "Distance in pixels from the top of the cell to the top of the strikethrough")
	cmd.Flags().String("adjust-strikethrough-thickness", "", "Thickness in pixels of the strikethrough")
	cmd.Flags().String("adjust-underline-position", "", "Distance in pixels from the top of the cell to the top of the underline")
	cmd.Flags().String("adjust-underline-thickness", "", "Thickness in pixels of the underline")
	cmd.Flags().String("alpha-blending", "", "What color space to use when performing alpha blending")
	cmd.Flags().String("app-notifications", "", "Control the in-app notifications that Ghostty shows")
	cmd.Flags().String("async-backend", "", "Configures the low-level API to use for async IO, eventing, etc.")
	cmd.Flags().String("auto-update", "", "Control the auto-update functionality of Ghostty")
	cmd.Flags().String("auto-update-channel", "", "The release channel to use for auto-updates")
	cmd.Flags().String("background", "", "Background color for the window")
	cmd.Flags().String("background-blur", "", "Whether to blur the background when `background-opacity` is less than 1")
	cmd.Flags().String("background-image", "", "Background image for the terminal")
	cmd.Flags().String("background-image-fit", "", "Background image fit")
	cmd.Flags().String("background-image-opacity", "", "Background image opacity")
	cmd.Flags().String("background-image-position", "", "Background image position")
	cmd.Flags().Bool("background-image-repeat", false, "Whether to repeat the background image or not")
	cmd.Flags().String("background-opacity", "", "The opacity level (opposite of transparency) of the background")
	cmd.Flags().Bool("background-opacity-cells", false, "Applies background opacity to cells with an explicit background color set")
	cmd.Flags().String("bell-audio-path", "", "If `audio` is an enabled bell feature, this is a path to an audio file")
	cmd.Flags().String("bell-audio-volume", "", "If `audio` is an enabled bell feature, this is the volume to play the audio file at")
	cmd.Flags().String("bell-features", "", "Bell features to enable if bell support is available in your runtime")
	cmd.Flags().String("bold-color", "", "Modifies the color used for bold text in the terminal")
	cmd.Flags().String("class", "", "The setting that will change the application class value")
	cmd.Flags().String("click-repeat-interval", "", "The time in milliseconds between clicks to consider a click a repeat")
	cmd.Flags().Bool("clipboard-paste-bracketed-safe", false, "If true, bracketed pastes will be considered safe")
	cmd.Flags().Bool("clipboard-paste-protection", false, "Require confirmation before pasting text that appears unsafe")
	cmd.Flags().String("clipboard-read", "", "Whether to allow programs running in the terminal to read from the system clipboard")
	cmd.Flags().Bool("clipboard-trim-trailing-spaces", false, "Trims trailing whitespace on data that is copied to the clipboard")
	cmd.Flags().String("clipboard-write", "", "Whether to allow programs running in the terminal to write to the system clipboard")
	cmd.Flags().String("command", "", "The command to run, usually a shell")
	cmd.Flags().StringArray("command-palette-entry", nil, "Custom entries into the command palette")
	cmd.Flags().Bool("config-default-files", false, "When this is true, the default configuration file paths will be loaded")
	cmd.Flags().StringArray("config-file", nil, "Additional configuration files to read")
	cmd.Flags().Bool("confirm-close-surface", false, "Confirms that a surface should be closed before closing it")
	cmd.Flags().String("copy-on-select", "", "Whether to automatically copy selected text to the clipboard")
	cmd.Flags().Bool("cursor-click-to-move", false, "Enables the ability to move the cursor at prompts")
	cmd.Flags().String("cursor-color", "", "The color of the cursor")
	cmd.Flags().String("cursor-opacity", "", "The opacity level (opposite of transparency) of the cursor")
	cmd.Flags().String("cursor-style", "", "The style of the cursor")
	cmd.Flags().Bool("cursor-style-blink", false, "Sets the default blinking state of the cursor")
	cmd.Flags().String("cursor-text", "", "The color of the text under the cursor")
	cmd.Flags().StringArray("custom-shader", nil, "Custom shaders to run after the default shaders")
	cmd.Flags().String("custom-shader-animation", "", "Controls when custom shaders are animated")
	cmd.Flags().Bool("desktop-notifications", false, "bool")
	cmd.Flags().String("enquiry-response", "", "String to send when we receive `ENQ` (`0x05`) from the command that we are running")
	cmd.Flags().StringArray("env", nil, "Extra environment variables to pass to commands launched in a terminal surface")
	cmd.Flags().String("faint-opacity", "", "The opacity level (opposite of transparency) of the faint text")
	cmd.Flags().Bool("focus-follows-mouse", false, "If true, when there are multiple split panes, the mouse selects the pane that is focused")
	cmd.Flags().StringArray("font-codepoint-map", nil, "Force one or a range of Unicode codepoints to map to a specific named font")
	cmd.Flags().StringArray("font-family", nil, "The font families to use")
	cmd.Flags().StringArray("font-family-bold", nil, "The font families to use for bold")
	cmd.Flags().StringArray("font-family-bold-italic", nil, "The font families to use for bold italic")
	cmd.Flags().StringArray("font-family-italic", nil, "The font families to use for italic")
	cmd.Flags().StringArray("font-feature", nil, "Apply a font feature")
	cmd.Flags().String("font-shaping-break", "", "Locations to break font shaping into multiple runs")
	cmd.Flags().String("font-size", "", "Font size in points")
	cmd.Flags().String("font-style", "", "The named font style to use")
	cmd.Flags().String("font-style-bold", "", "The named font style to use for bold")
	cmd.Flags().String("font-style-bold-italic", "", "The named font style to use for bold italic")
	cmd.Flags().String("font-style-italic", "", "The named font style to use for italic")
	cmd.Flags().String("font-synthetic-style", "", "Control whether Ghostty should synthesize a style if the requested style is not available")
	cmd.Flags().Bool("font-thicken", false, "Draw fonts with a thicker stroke")
	cmd.Flags().String("font-thicken-strength", "", "Strength of thickening when `font-thicken` is enabled")
	cmd.Flags().StringArray("font-variation", nil, "Set font variation")
	cmd.Flags().StringArray("font-variation-bold", nil, "Set font variation for bold")
	cmd.Flags().StringArray("font-variation-bold-italic", nil, "Set font variation for bold-italic")
	cmd.Flags().StringArray("font-variation-italic", nil, "Set font variation for italic")
	cmd.Flags().String("foreground", "", "Foreground color for the window")
	cmd.Flags().String("freetype-load-flags", "", "FreeType load flags to enable")
	cmd.Flags().Bool("fullscreen", false, "Start new windows in fullscreen")
	cmd.Flags().String("grapheme-width-method", "", "The method to use for calculating the cell width of a grapheme cluster")
	cmd.Flags().String("gtk-single-instance", "", "Run in single-instance mode")
	cmd.Flags().String("gtk-tabs-location", "", "Determines the side of the screen that the GTK tab bar will stick to")
	cmd.Flags().Bool("gtk-titlebar", false, "Display the full GTK titlebar")
	cmd.Flags().Bool("gtk-toolbar-style", false, "Determines the appearance of the top and bottom bars tab bar")
	cmd.Flags().Bool("gtk-wide-tabs", false, "Use \"wide\" GTK tabs")
	cmd.Flags().String("image-storage-limit", "", "The total amount of bytes that can be used for image data")
	cmd.Flags().String("initial-command", "", "This is the same as \"command\", but only applies to the first terminal surface created")
	cmd.Flags().Bool("initial-window", false, "Create an initial window when Ghostty is run")
	cmd.Flags().String("input", "", "Data to send as input to the command on startup")
	cmd.Flags().Bool("link-previews", false, "Show link previews for a matched URL")
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
	cmd.Flags().String("quick-terminal-keyboard-interactivity", "", "Determines under which circumstances that the quick terminal should receive keyboard input")
	cmd.Flags().String("quick-terminal-position", "", "The position of the \"quick\" terminal window")
	cmd.Flags().String("quick-terminal-screen", "", "The screen where the quick terminal should show up")
	cmd.Flags().Bool("quit-after-last-window-closed", false, "Whether or not to quit after the last surface is closed")
	cmd.Flags().String("quit-after-last-window-closed-delay", "", "Controls how long Ghostty will stay running after the last open surface has been closed")
	cmd.Flags().String("resize-overlay", "", "Controls when resize overlays are shown")
	cmd.Flags().String("resize-overlay-duration", "", "Controls how long the overlay is visible on the screen before it is hidde")
	cmd.Flags().String("resize-overlay-position", "", "Controls the position of the overlay")
	cmd.Flags().String("scroll-to-bottom", "", "When to scroll the surface to the bottom")
	cmd.Flags().String("scrollback-limit", "", "The size of the scrollback buffer in bytes")
	cmd.Flags().String("selection-background", "", "The background color for selection")
	cmd.Flags().String("selection-foreground", "", "The foreground color for selection")
	cmd.Flags().String("shell-integration", "", "Whether to enable shell integration auto-injection or not")
	cmd.Flags().String("shell-integration-features", "", "Shell integration features to enable")
	cmd.Flags().String("split-divider-color", "", "The color of the split divider")
	cmd.Flags().String("theme", "", "A theme to use")
	cmd.Flags().String("title", "", "The title Ghostty will use for the window")
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
	cmd.Flags().String("window-position-x", "", "The starting window position")
	cmd.Flags().String("window-position-y", "", "The starting window position")
	cmd.Flags().String("window-save-state", "", "Whether to enable saving and restoring window state")
	cmd.Flags().Bool("window-step-resize", false, "Resize the window in discrete increments of the focused surface's cell size")
	cmd.Flags().String("window-theme", "", "The theme to use for the windows")
	cmd.Flags().String("window-title-font-family", "", "The font that will be used for the application's window and tab titles")
	cmd.Flags().Bool("window-vsync", false, "Synchronize rendering with the screen refresh rate")
	cmd.Flags().String("window-width", "", "The initial window width")
	cmd.Flags().String("working-directory", "", "The directory to change to after starting the command")
	cmd.Flags().String("x11-instance-name", "", "This controls the instance name field of the `WM_CLASS` X11 property")

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Value.Type() != "bool" {
			f.NoOptDefVal = " " // flag arguments need to be passed in attached mode
		}
	})

	// TODO use font-families from `+list-fonts`?
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"alpha-blending":      ghostty.ActionAlphaBlendings(),
		"app-notifications":   ghostty.ActionNotifications().UniqueList(","),
		"async-backend":       ghostty.ActionAsyncBackends(),
		"auto-update":         ghostty.ActionAutoUpdateModes(),
		"auto-update-channel": ghostty.ActionReleaseChannels(),
		"background": carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
		).ToA(),
		"background-blur": carapace.ActionValuesDescribed(
			"false", "equivalent to a blur intensity of 0",
			"true", "equivalent to the default blur intensity of 20",
		).StyleF(style.ForKeyword),
		"background-image":          carapace.ActionFiles(),
		"background-image-fit":      ghostty.ActionBackgroundImageFits(),
		"background-image-position": ghostty.ActionBackgroundImagePositions(),
		"bell-audio-path":           carapace.ActionFiles(),
		"bell-features":             ghostty.ActionBellFeatures().UniqueList(","),
		"bold-color": carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
			carapace.ActionValues("bright"),
		).ToA(),
		"clipboard-read":          carapace.ActionValues("ask", "allow", "deny").StyleF(style.ForKeyword),
		"clipboard-write":         carapace.ActionValues("ask", "allow", "deny").StyleF(style.ForKeyword),
		"command":                 bridge.ActionCarapaceBin().SplitP(),
		"config-file":             carapace.ActionFiles(),
		"copy-on-select":          ghostty.ActionCopyOnSelectModes(),
		"cursor-color":            color.ActionHexColors(),
		"cursor-style":            ghostty.ActionCursorStyles(),
		"cursor-text":             color.ActionHexColors(),
		"custom-shader":           carapace.ActionFiles(),
		"custom-shader-animation": ghostty.ActionShaderAnimationModes(),
		"env":                     env.ActionNameValues(false),
		"font-codepoint-map": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues() // TODO codepoint
			default:
				return ghostty.ActionFontFamilies()
			}
		}),
		"font-family":             ghostty.ActionFontFamilies(),
		"font-family-bold":        ghostty.ActionFontFamilies(),
		"font-family-bold-italic": ghostty.ActionFontFamilies(),
		"font-family-italic":      ghostty.ActionFontFamilies(),
		"font-feature":            carapace.ActionValues(), // TODO font-features
		"font-shaping-break": carapace.ActionValuesDescribed(
			"cursor", "Break runs under the cursor",
		),
		"font-style":                 carapace.ActionValues(), // TODO font style
		"font-style-bold":            carapace.ActionValues(), // TODO font style
		"font-style-bold-italic":     carapace.ActionValues(), // TODO font style
		"font-style-italic":          carapace.ActionValues(), // TODO font style
		"font-synthetic-style":       ghostty.ActionFontSyntheticStyles().UniqueList(","),
		"font-variation":             carapace.ActionValues(), // TODO font variation
		"font-variation-bold":        carapace.ActionValues(), // TODO font variation
		"font-variation-bold-italic": carapace.ActionValues(), // TODO font variation
		"font-variation-italic":      carapace.ActionValues(), // TODO font variation
		"foreground": carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
		).ToA(),
		"freetype-load-flags":   ghostty.ActionFreetypeLoadFlags().UniqueList(","),
		"grapheme-width-method": ghostty.ActionGraphemeWidthMethods(),
		"gtk-single-instance":   carapace.ActionValues("true", "false", "detect").StyleF(style.ForKeyword),
		"gtk-tabs-location":     carapace.ActionValues("top", "bottom", "left", "right", "hidden"),
		"gtk-toolbar-style":     ghostty.ActionGtkToolbarStyles(),
		"initial-command":       bridge.ActionCarapaceBin().Split(),
		"input": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"raw", "Send raw text as-is",
					"path", "Read a filepath and send the contents",
				).Suffix(":")
			default:
				switch c.Parts[0] {
				case "path":
					return carapace.ActionFiles()
				default:
					return carapace.ActionValues()
				}
			}
		}),
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
		"quick-terminal-keyboard-interactivity": carapace.ActionValuesDescribed(
			"none", "The quick terminal will not receive any keyboard input",
			"on-demand", "The quick terminal would only receive keyboard input when it is focused",
			"exclusive", "The quick terminal will always receive keyboard input",
		).StyleF(style.ForKeyword),
		"quick-terminal-position": ghostty.ActionQuickTerminalPositions(),
		"quick-terminal-screen":   ghostty.ActionQuickTerminalScreens(),
		"resize-overlay":          ghostty.ActionResizeOverlayModes(),
		"resize-overlay-position": ghostty.ActionResizeOverlayPositions(),
		"scroll-to-bottom": carapace.ActionValuesDescribed(
			"keystroke", "scroll the surface to the bottom when the user presses a key",
			"output", "scroll the surface to the bottom if there is new data to display",
		).UniqueList(","),
		"selection-background":       color.ActionHexColors(),
		"selection-foreground":       color.ActionHexColors(),
		"shell-integration":          ghostty.ActionShellIntegrationModes(),
		"shell-integration-features": ghostty.ActionShellIntegrationFeatures().UniqueList(","),
		"split-divider-color": carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
		).ToA(),
		"theme":                    ghostty.ActionThemes(),
		"unfocused-split-fill":     color.ActionXtermColorNames(),
		"window-colorspace":        carapace.ActionValues("srgb", "display-p3"),
		"window-new-tab-position":  ghostty.ActionWindowNewTabPositions(),
		"window-padding-color":     ghostty.ActionWindowPaddingColors(),
		"window-save-state":        ghostty.ActionWindowSaveStates(),
		"window-theme":             ghostty.ActionWindowThemes(),
		"window-title-font-family": ghostty.ActionFontFamilies(),
		"working-directory":        carapace.ActionDirectories(),
	})
}
