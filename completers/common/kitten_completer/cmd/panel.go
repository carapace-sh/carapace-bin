package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panelCmd = &cobra.Command{
	Use:   "panel",
	Short: "Use a command line program to draw a GPU accelerated panel on your desktop",
}

func init() {
	rootCmd.AddCommand(panelCmd)
	carapace.Gen(panelCmd).Standalone()

	panelCmd.Flags().Float64("lines", 0, "The number of lines shown in the panel. Ignored for background, centered, and vertical panels. If it has the suffix px then it sets the height of the panel in pixels instead of lines.")
	panelCmd.Flags().Float64("columns", 0, "The number of columns shown in the panel. Ignored for background, centered, and horizontal panels. If it has the suffix px then it sets the width of the panel in pixels instead of columns.")
	panelCmd.Flags().Float64("margin-top", 0, "Set the top margin for the panel, in pixels. Has no effect for bottom edge panels. Only works on macOS and Wayland compositors that supports the wlr layer shell protocol.")
	panelCmd.Flags().Float64("margin-left", 0, "Set the left margin for the panel, in pixels. Has no effect for right edge panels. Only works on macOS and Wayland compositors that supports the wlr layer shell protocol.")
	panelCmd.Flags().Float64("margin-bottom", 0, "Set the bottom margin for the panel, in pixels. Has no effect for top edge panels. Only works on macOS and Wayland compositors that supports the wlr layer shell protocol.")
	panelCmd.Flags().Float64("margin-right", 0, "Set the right margin for the panel, in pixels. Has no effect for left edge panels. Only works on macOS and Wayland compositors that supports the wlr layer shell protocol.")
	panelCmd.Flags().String("edge", "", "Which edge of the screen to place the panel on. Note that some window managers (such as i3) do not support placing docked windows on the left and right edges. The value background means make the panel the \"desktop wallpaper\". Note that when using sway if you set a background in your sway config it will cover the background drawn using this kitten. Additionally, there are three more values: center, center-sized and none. The value center anchors the panel to all sides and covers the entire display (on macOS the part of the display not covered by titlebar and dock). The panel can be shrunk and placed using the margin parameters. The value none anchors the panel to the top left corner and should be placed using the margin parameters. Its size is set by --lines and --columns. The value center-sized is just like none except that the panel is centered instead of in the top left corner and the margins have no effect.")
	panelCmd.Flags().String("layer", "", "On a Wayland compositor that supports the wlr layer shell protocol, specifies the layer on which the panel should be drawn. This parameter is ignored and set to background if --edge is set to background. On macOS, maps these to appropriate NSWindow *levels*.")
	panelCmd.Flags().StringP("config", "c", "", "Path to config file to use for kitty when drawing the panel.")
	panelCmd.Flags().StringArrayP("override", "o", nil, "Override individual kitty configuration options, can be specified multiple times. Syntax: name=value. For example: -o font_size=20")
	panelCmd.Flags().String("output-name", "", "The panel can only be displayed on a single monitor (output) at a time. This allows you to specify which output is used, by name. If not specified the compositor will choose an output automatically, typically the last output the user interacted with or the primary monitor. Use the special value list to get a list of available outputs. Use listjson for a json encoded output. Note that on Wayland the output can only be set at panel creation time, it cannot be changed after creation, nor is there anyway to display a single panel on all outputs. Please complain to the Wayland developers about this.")
	panelCmd.Flags().String("focus-policy", "", "On a Wayland compositor that supports the wlr layer shell protocol, specify the focus policy for keyboard interactivity with the panel. Please refer to the wlr layer shell protocol documentation for more details. Note that different Wayland compositors behave very differently with exclusive, your mileage may vary. On macOS, exclusive and on-demand are currently the same.")
	panelCmd.Flags().Bool("hide-on-focus-loss", false, "Automatically hide the panel window when it loses focus. Using this option will force --focus-policy to on-demand. Note that on Wayland, depending on the compositor, this can result in the window never becoming visible.")
	panelCmd.Flags().Bool("grab-keyboard", false, "Grab the keyboard. This means global shortcuts defined in the OS will be passed to kitty instead. Useful if you want to create an OS modal window. How well this works depends on the OS/window manager/desktop environment. On Wayland it works only if the compositor implements the inhibit-keyboard-shortcuts protocol. On macOS Apple doesn't allow applications to grab the keyboard without special permissions, so it doesn't work.")
	panelCmd.Flags().Float64("exclusive-zone", 0, "On a Wayland compositor that supports the wlr layer shell protocol, request a given exclusive zone for the panel. Please refer to the wlr layer shell documentation for more details on the meaning of exclusive and its value. If --edge is set to anything other than center or none, this flag will not have any effect unless the flag --override-exclusive-zone is also set. If --edge is set to background, this option has no effect. Ignored on X11 and macOS.")
	panelCmd.Flags().Bool("override-exclusive-zone", false, "On a Wayland compositor that supports the wlr layer shell protocol, override the default exclusive zone. This has effect only if --edge is set to top, left, bottom or right. Ignored on X11 and macOS.")
	panelCmd.Flags().BoolP("single-instance", "1", false, "If specified only a single instance of the panel will run. New invocations will instead create a new top-level window in the existing panel instance.")
	panelCmd.Flags().String("instance-group", "", "Used in combination with the --single-instance option. All panel invocations with the same --instance-group will result in new panels being created in the first panel instance within that group.")
	panelCmd.Flags().Bool("wait-for-single-instance-window-close", false, "Normally, when using --single-instance, kitty will open a new window in an existing instance and quit immediately. With this option, it will not quit till the newly opened window is closed. Note that if no previous instance is found, then kitty will wait anyway, regardless of this option.")
	panelCmd.Flags().String("listen-on", "", "Listen on the specified socket address for control messages. For example, --listen-on=unix:/tmp/mykitty or --listen-on=tcp:localhost:12345. On Linux systems, you can also use abstract UNIX sockets, not associated with a file, like this: --listen-on=unix:@mykitty. Environment variables are expanded and relative paths are resolved with respect to the temporary directory. To control kitty, you can send commands to it with kitten @ using the --to option to specify this address. Note that if you run kitten @ within a kitty window, there is no need to specify the --to option as it will automatically read from the environment. Note that this will be ignored unless allow_remote_control is set to either: yes, socket or socket-only. This can also be specified in kitty.conf.")
	panelCmd.Flags().Bool("toggle-visibility", false, "When set and using --single-instance will toggle the visibility of the existing panel rather than creating a new one.")
	panelCmd.Flags().Bool("move-to-active-monitor", false, "When set and using --toggle-visibility to show an existing panel, the panel is moved to the active monitor (typically the monitor with the mouse on it). This works only if the underlying OS supports it. It is currently supported on macOS only.")
	panelCmd.Flags().Bool("start-as-hidden", false, "Start in hidden mode, useful with --toggle-visibility.")
	panelCmd.Flags().Bool("detach", false, "Detach from the controlling terminal, if any, running in an independent child process, the parent process exits immediately.")
	panelCmd.Flags().String("detached-log", "", "Path to a log file to store STDOUT/STDERR when using --detach")
	panelCmd.Flags().Bool("debug-rendering", false, "For internal debugging use.")
	panelCmd.Flags().Bool("debug-input", false, "For internal debugging use.")
	panelCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(panelCmd).FlagCompletion(carapace.ActionMap{
		"edge":         carapace.ActionValues("top", "background", "bottom", "center", "center-sized", "left", "none", "right"),
		"layer":        carapace.ActionValues("bottom", "background", "overlay", "top"),
		"focus-policy": carapace.ActionValues("not-allowed", "exclusive", "on-demand"),
		"listen-on":    carapace.ActionValues("unix:", "tcp:").NoSpace(':'),
	})

	carapace.Gen(panelCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
