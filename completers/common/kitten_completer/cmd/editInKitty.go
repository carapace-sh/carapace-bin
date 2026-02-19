package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editInKittyCmd = &cobra.Command{
	Use:   "edit-in-kitty",
	Short: "Edit a file in a kitty overlay window",
}

func init() {
	rootCmd.AddCommand(editInKittyCmd)
	carapace.Gen(editInKittyCmd).Standalone()

	editInKittyCmd.Flags().String("title", "", "The title to set for the new window. By default, title is controlled by the child process. The special value current will copy the title from the --source-window.")
	editInKittyCmd.Flags().String("window-title", "", "The title to set for the new window. By default, title is controlled by the child process. The special value current will copy the title from the --source-window.")
	editInKittyCmd.Flags().String("tab-title", "", "The title for the new tab if launching in a new tab. By default, the title of the active window in the tab is used as the tab title. The special value current will copy the title from the tab containing the --source-window.")
	editInKittyCmd.Flags().String("type", "", "Where to launch the child process:")
	editInKittyCmd.Flags().Bool("dont-take-focus", false, "Keep the focus on the currently active window instead of switching to the newly opened window.")
	editInKittyCmd.Flags().Bool("keep-focus", false, "Keep the focus on the currently active window instead of switching to the newly opened window.")
	editInKittyCmd.Flags().String("cwd", "", "The working directory for the newly launched child. Use the special value current to use the working directory of the --source-window. The special value last_reported uses the last working directory reported by the shell (needs shell_integration to work). The special value oldest works like current but uses the working directory of the oldest foreground process associated with the currently active window rather than the newest foreground process. Finally, the special value root refers to the process that was originally started when the window was created. When opening in the same working directory as the current window causes the new window to connect to a remote host, you can use the --hold-after-ssh flag to prevent the new window from closing when the connection is terminated.")
	editInKittyCmd.Flags().StringArray("env", nil, "Environment variables to set in the child process. Can be specified multiple times to set different environment variables. Syntax: name=value. Using name= will set to empty string and just name will remove the environment variable.")
	editInKittyCmd.Flags().StringArray("var", nil, "User variables to set in the created window. Can be specified multiple times to set different user variables. Syntax: name=value. Using name= will set to empty string.")
	editInKittyCmd.Flags().Bool("hold", false, "Keep the window open even after the command being executed exits, at a shell prompt. The shell will be run after the launched command exits.")
	editInKittyCmd.Flags().String("location", "", "Where to place the newly created window when it is added to a tab which already has existing windows in it. after and before place the new window before or after the active window. neighbor is a synonym for after. Also applies to creating a new tab, where the value of after will cause the new tab to be placed next to the current tab instead of at the end. The values of vsplit, hsplit and split are only used by the splits layout and control if the new window is placed in a vertical, horizontal or automatic split with the currently active window. The default is to place the window in a layout dependent manner, typically, after the currently active window. See --next-to to use a window other than the currently active window.")
	editInKittyCmd.Flags().String("next-to", "", "A match expression to select the window next to which the new window is created. See search_syntax for the syntax for specifying windows. If not specified defaults to the active window. When used via remote control and a target tab is specified this option is ignored unless the matched window is in the specified tab. When using --type of tab, the tab will be created in the OS Window containing the matched window.")
	editInKittyCmd.Flags().String("os-window-class", "", "Set the WM_CLASS property on X11 and the application id property on Wayland for the newly created OS window when using --type. Defaults to whatever is used by the parent kitty process, which in turn defaults to kitty.")
	editInKittyCmd.Flags().String("os-window-name", "", "Set the WM_NAME property on X11 for the newly created OS Window when using --type. Defaults to --os-window-class.")
	editInKittyCmd.Flags().String("os-window-title", "", "Set the title for the newly created OS window. This title will override any titles set by programs running in kitty. The special value current will copy the title from the OS Window containing the --source-window.")
	editInKittyCmd.Flags().String("os-window-state", "", "The initial state for the newly created OS Window.")
	editInKittyCmd.Flags().String("logo", "", "Path to a PNG image to use as the logo for the newly created window. See window_logo_path. Relative paths are resolved from the kitty configuration directory.")
	editInKittyCmd.Flags().String("logo-position", "", "The position for the window logo. Only takes effect if --logo is specified. See window_logo_position.")
	editInKittyCmd.Flags().Float64("logo-alpha", 0, "The amount the window logo should be faded into the background. Only takes effect if --logo is specified. See window_logo_alpha.")
	editInKittyCmd.Flags().String("color", "", "Change colors in the newly launched window. You can either specify a path to a .conf file with the same syntax as kitty.conf to read the colors from, or specify them individually, for example:: --color background=white --color foreground=red")
	editInKittyCmd.Flags().String("spacing", "", "Set the margin and padding for the newly created window. For example: margin=20 or padding-left=10 or margin-h=30. The shorthand form sets all values, the *-h and *-v variants set horizontal and vertical values. Can be specified multiple times. Note that this is ignored for overlay windows as these use the settings from the base window.")
	editInKittyCmd.Flags().Bool("hold-after-ssh", false, "When using --cwd=current or similar from a window that is running the ssh kitten, the new window will run a local shell after disconnecting from the remote host, when this option is specified.")
	editInKittyCmd.Flags().Float64("max-file-size", 0, "The maximum allowed size (in MB) of files to edit. Since the file data has to be base64 encoded and transmitted over the tty device, overly large files will not perform well.")
	editInKittyCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(editInKittyCmd).FlagCompletion(carapace.ActionMap{
		"title":        carapace.ActionValues("current"),
		"window-title": carapace.ActionValues("current"),
		"tab-title":    carapace.ActionValues("current"),
		"type": carapace.ActionValuesDescribed(
			"window", "A new kitty window in the current tab",
			"tab", "A new tab in the current OS window. Not available when the The launch command command is used in startup sessions.",
			"os-window", "A new operating system window.  Not available when the The launch command command is used in startup sessions.",
			"overlay", "An overlay window covering the current active kitty window",
			"overlay-main", "An overlay window covering the current active kitty window. Unlike a plain overlay window, this window is considered as a main window which means it is used as the active window for getting the current working directory, the input text for kittens, launch commands, etc. Useful if this overlay is intended to run for a long time as a primary window.",
			"background", "The process will be run in the background, without a kitty window. Note that if --allow-remote-control is specified the KITTY_LISTEN_ON environment variable will be set to a dedicated socket pair file descriptor that the process can use for remote control.",
			"clipboard", "These two are meant to work with --stdin-source to copy data to the system clipboard or primary selection.",
			"primary", "These two are meant to work with --stdin-source to copy data to the system clipboard or primary selection.",
			"os-panel", "Similar to os-window, except that it creates the new OS Window as a desktop panel. Only works on platforms that support this, such as Wayand compositors that support the layer shell protocol. Use the --os-panel option to configure the panel.",
		),
		"cwd":             carapace.ActionValues("current", "last-reported", "oldest", "root"),
		"location":        carapace.ActionValues("default", "after", "before", "first", "hsplit", "last", "neighbor", "split", "vsplit"),
		"os-window-state": carapace.ActionValues("normal", "fullscreen", "maximized", "minimized"),
		"logo":            carapace.ActionFiles("png").Chdir("~/.config/kitty"),
		"color":           carapace.ActionFiles("conf"),
	})

	carapace.Gen(editInKittyCmd).PositionalCompletion(carapace.ActionFiles())
}
