package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Change the behaviour of zellij",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optionsCmd).Standalone()

	optionsCmd.Flags().String("advanced-mouse-actions", "", "Whether to enable mouse hover effects and pane grouping functionality default is true")
	optionsCmd.Flags().String("attach-to-session", "", "Whether to attach to a session specified in \"session-name\" if it exists")
	optionsCmd.Flags().String("auto-layout", "", "Whether to lay out panes in a predefined set of layouts whenever possible")
	optionsCmd.Flags().String("client-async-worker-tasks", "", "Number of async worker tasks to spawn per active client")
	optionsCmd.Flags().String("copy-clipboard", "", "OSC52 destination clipboard")
	optionsCmd.Flags().String("copy-command", "", "Switch to using a user supplied command for clipboard instead of OSC52")
	optionsCmd.Flags().String("copy-on-select", "", "Automatically copy when selecting text (true or false)")
	optionsCmd.Flags().String("dangerously-enable-paste-buffer-read", "", "")
	optionsCmd.Flags().String("default-cwd", "", "Set the default cwd")
	optionsCmd.Flags().String("default-layout", "", "Set the default layout")
	optionsCmd.Flags().String("default-mode", "", "Set the default mode")
	optionsCmd.Flags().String("default-shell", "", "Set the default shell")
	optionsCmd.Flags().String("disable-session-metadata", "", "If true, will disable writing session metadata to disk")
	optionsCmd.Flags().String("focus-follows-mouse", "", "Whether to focus panes on mouse hover (true or false) default is false")
	optionsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	optionsCmd.Flags().String("layout-dir", "", "Set the layout_dir, defaults to subdirectory of config dir")
	optionsCmd.Flags().String("mirror-session", "", "Mirror session when multiple users are connected (true or false)")
	optionsCmd.Flags().String("mouse-click-through", "", "Whether clicking a pane to focus it also sends the click into the pane (true or false) default is false")
	optionsCmd.Flags().String("mouse-hover-effects", "", "Whether to enable mouse hover visual effects (frame highlight and help text) default is true")
	optionsCmd.Flags().String("mouse-mode", "", "Set the handling of mouse events (true or false) Can be temporarily bypassed by the [SHIFT] key")
	optionsCmd.Flags().String("mouse-scroll-resize", "", "Whether Ctrl+ScrollWheel resizes panes default is true")
	optionsCmd.Flags().String("nested-session-handling", "", "How to handle a nested Zellij session detected inside a pane (ask, fullscreen, descend, never)")
	optionsCmd.Flags().String("on-force-close", "", "Set behaviour on force close (quit or detach)")
	optionsCmd.Flags().String("osc133-command-selection", "", "Whether triple-clicking inside shell-marked (OSC 133) command output selects the command and its output rather than the logical line default is true")
	optionsCmd.Flags().String("osc8-hyperlinks", "", "Enable OSC8 hyperlink output (true or false)")
	optionsCmd.Flags().String("pane-frame-style", "", "")
	optionsCmd.Flags().String("pane-frames", "", "Set display of the pane frames (true or false)")
	optionsCmd.Flags().String("post-command-discovery-hook", "", "A command to run after the discovery of running commands when serializing, for the purpose of manipulating the command (eg. with a regex) before it gets serialized")
	optionsCmd.Flags().String("scroll-buffer-size", "", "")
	optionsCmd.Flags().String("scrollback-editor", "", "Explicit full path to open the scrollback editor (default is $EDITOR or $VISUAL)")
	optionsCmd.Flags().String("scrollback-lines-to-serialize", "", "Scrollback lines to serialize along with the pane viewport when serializing sessions, 0 defaults to the scrollback size. If this number is higher than the scrollback size, it will also default to the scrollback size")
	optionsCmd.Flags().String("serialization-interval", "", "The interval at which to serialize sessions for resurrection (in seconds)")
	optionsCmd.Flags().String("serialize-pane-viewport", "", "Whether pane viewports are serialized along with the session, default is false")
	optionsCmd.Flags().String("session-name", "", "The name of the session to create when starting Zellij")
	optionsCmd.Flags().String("session-serialization", "", "Whether sessions should be serialized to the HD so that they can be later resurrected, default is true")
	optionsCmd.Flags().String("show-release-notes", "", "Whether to show release notes on first run of a new version default is true")
	optionsCmd.Flags().String("show-startup-tips", "", "Whether to show startup tips when starting a new session default is true")
	optionsCmd.Flags().String("simplified-ui", "", "Allow plugins to use a more simplified layout that is compatible with more fonts (true or false)")
	optionsCmd.Flags().String("stacked-pane-list", "", "")
	optionsCmd.Flags().String("stacked-resize", "", "Whether to stack panes when resizing beyond a certain size default is true")
	optionsCmd.Flags().String("styled-underlines", "", "Whether to use ANSI styled underlines")
	optionsCmd.Flags().String("support-kitty-graphics-protocol", "", "Whether to enable support for the Kitty graphics (image) protocol (must also be supported by the host terminal), defaults to true if the terminal supports it")
	optionsCmd.Flags().String("support-kitty-keyboard-protocol", "", "Whether to enable support for the Kitty keyboard protocol (must also be supported by the host terminal), defaults to true if the terminal supports it")
	optionsCmd.Flags().String("theme", "", "Set the default theme")
	optionsCmd.Flags().String("theme-dark", "", "Theme name to apply when the host terminal reports a dark color palette (CSI 2031 / DSR 997). Requires `theme_light` to also be set; if either is missing the static `theme` remains authoritative")
	optionsCmd.Flags().String("theme-dir", "", "Set the theme_dir, defaults to subdirectory of config dir")
	optionsCmd.Flags().String("theme-light", "", "Theme name to apply when the host terminal reports a light color palette (CSI 2031 / DSR 997). Requires `theme_dark` to also be set; if either is missing the static `theme` remains authoritative")
	optionsCmd.Flags().String("visual-bell", "", "Whether to show visual bell indicators (pane/tab frame flash and [!] suffix) default is true")
	optionsCmd.Flags().String("web-server", "", "Whether to make sure a local web server is running when a new Zellij session starts. This web server will allow creating new sessions and attaching to existing ones that have opted in to being shared in the browser")
	optionsCmd.Flags().String("web-sharing", "", "Whether to allow new sessions to be shared through a local web server, assuming one is running (see the `web_server` option for more details)")
	optionsCmd.Flags().String("word-separators", "", "Characters that terminate a word when double-clicking to select it, in addition to whitespace (which is always a separator) default is \"[]{}<>()\"")
	rootCmd.AddCommand(optionsCmd)

	carapace.Gen(optionsCmd).FlagCompletion(carapace.ActionMap{
		"advanced-mouse-actions":               carapace.ActionValues("true", "false"),
		"attach-to-session":                    carapace.ActionValues("true", "false"),
		"auto-layout":                          carapace.ActionValues("true", "false"),
		"copy-clipboard":                       carapace.ActionValues("system", "primary"),
		"copy-on-select":                       carapace.ActionValues("true", "false"),
		"dangerously-enable-paste-buffer-read": carapace.ActionValues("true", "false"),
		"default-cwd":                          carapace.ActionFiles(),
		"default-layout":                       carapace.ActionFiles(),
		"default-mode":                         actionModes(),
		"default-shell":                        carapace.ActionFiles(),
		"disable-session-metadata":             carapace.ActionValues("true", "false"),
		"focus-follows-mouse":                  carapace.ActionValues("true", "false"),
		"layout-dir":                           carapace.ActionFiles(),
		"mirror-session":                       carapace.ActionValues("true", "false"),
		"mouse-click-through":                  carapace.ActionValues("true", "false"),
		"mouse-hover-effects":                  carapace.ActionValues("true", "false"),
		"mouse-mode":                           carapace.ActionValues("true", "false"),
		"mouse-scroll-resize":                  carapace.ActionValues("true", "false"),
		"nested-session-handling":              carapace.ActionValues("ask", "fullscreen", "descend", "never"),
		"on-force-close":                       carapace.ActionValues("quit", "detach"),
		"osc133-command-selection":             carapace.ActionValues("true", "false"),
		"osc8-hyperlinks":                      carapace.ActionValues("true", "false"),
		"pane-frame-style":                     carapace.ActionValues("full", "titles", "none"),
		"pane-frames":                          carapace.ActionValues("true", "false"),
		"scrollback-editor":                    carapace.ActionFiles(),
		"serialize-pane-viewport":              carapace.ActionValues("true", "false"),
		"session-serialization":                carapace.ActionValues("true", "false"),
		"show-release-notes":                   carapace.ActionValues("true", "false"),
		"show-startup-tips":                    carapace.ActionValues("true", "false"),
		"simplified-ui":                        carapace.ActionValues("true", "false"),
		"stacked-pane-list":                    carapace.ActionValues("true", "false"),
		"stacked-resize":                       carapace.ActionValues("true", "false"),
		"styled-underlines":                    carapace.ActionValues("true", "false"),
		"support-kitty-graphics-protocol":      carapace.ActionValues("true", "false"),
		"support-kitty-keyboard-protocol":      carapace.ActionValues("true", "false"),
		"theme-dir":                            carapace.ActionFiles(),
		"visual-bell":                          carapace.ActionValues("true", "false"),
		"web-server":                           carapace.ActionValues("true", "false"),
		"web-sharing":                          carapace.ActionValues("on", "off", "disabled"),
	})

	carapace.Gen(optionsCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionValues("true", "false"),
	)
}
