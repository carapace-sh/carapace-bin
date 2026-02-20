package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Run an arbitrary process in a new window/tab",
}

func init() {
	launchCmd.AddCommand(atCmd)
	carapace.Gen(launchCmd).Standalone()

	launchCmd.Flags().String("add-to-session", "", "Add the newly created window/tab to the specified session")
	launchCmd.Flags().Bool("allow-remote-control", false, "Programs running in this window can control kitty")
	launchCmd.Flags().String("bias", "0", "The bias used to alter the size of the window")
	launchCmd.Flags().String("color", "", "Change colors in the newly launched window")
	launchCmd.Flags().Bool("copy-cmdline", false, "Ignore any specified command line and instead use the command line from the source window")
	launchCmd.Flags().Bool("copy-colors", false, "Set the colors of the newly created window to be the same as the source window")
	launchCmd.Flags().Bool("copy-env", false, "Copy the environment variables from the source window into the newly launched child process")
	launchCmd.Flags().String("cwd", "", "The working directory for the newly launched child")
	launchCmd.Flags().Bool("dont-take-focus", false, "Keep the focus on the currently active window instead of switching to the newly opened window")
	launchCmd.Flags().String("env", "", "Environment variables to set in the child process")
	launchCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	launchCmd.Flags().Bool("hold", false, "Keep the window open even after the command being executed exits")
	launchCmd.Flags().Bool("hold-after-ssh", false, "Run a local shell after disconnecting from the remote host")
	launchCmd.Flags().String("location", "default", "Where to place the newly created window")
	launchCmd.Flags().String("logo", "", "Path to a PNG image to use as the logo for the newly created window")
	launchCmd.Flags().String("logo-alpha", "-1", "The amount the window logo should be faded into the background")
	launchCmd.Flags().String("logo-position", "", "The position for the window logo")
	launchCmd.Flags().String("marker", "", "Create a marker that highlights text in the newly created window")
	launchCmd.Flags().StringP("match", "m", "", "The tab to match")
	launchCmd.Flags().String("next-to", "", "A match expression to select the window next to which the new window is created")
	launchCmd.Flags().Bool("no-response", false, "Do not print out the id of the newly created window")
	launchCmd.Flags().String("os-panel", "", "Options to control the creation of desktop panels")
	launchCmd.Flags().String("os-window-class", "", "Set the WM_CLASS property on X11 and the application id property on Wayland")
	launchCmd.Flags().String("os-window-name", "", "Set the WM_NAME property on X11 for the newly created OS Window")
	launchCmd.Flags().String("os-window-state", "normal", "The initial state for the newly created OS Window")
	launchCmd.Flags().String("os-window-title", "", "Set the title for the newly created OS window")
	launchCmd.Flags().String("remote-control-password", "", "Restrict the actions remote control is allowed to take")
	launchCmd.Flags().String("response-timeout", "86400", "The time in seconds to wait for the started process to exit")
	launchCmd.Flags().Bool("self", false, "Use the tab containing the window this command is run in instead of the active tab")
	launchCmd.Flags().String("source-window", "", "A match expression to select the window from which data is copied")
	launchCmd.Flags().String("spacing", "", "Set the margin and padding for the newly created window")
	launchCmd.Flags().Bool("stdin-add-formatting", false, "When using --stdin-source add formatting escape codes")
	launchCmd.Flags().Bool("stdin-add-line-wrap-markers", false, "When using --stdin-source add a carriage return at every line wrap location")
	launchCmd.Flags().String("stdin-source", "none", "Pass the screen contents as STDIN to the child process")
	launchCmd.Flags().String("tab-title", "", "The title for the new tab if launching in a new tab")
	launchCmd.Flags().String("title", "", "The title to set for the new window")
	launchCmd.Flags().String("type", "window", "Where to launch the child process")
	launchCmd.Flags().String("var", "", "User variables to set in the created window")
	launchCmd.Flags().Bool("wait-for-child-to-exit", false, "Wait until the launched program exits and print out its exit code")
	launchCmd.Flags().StringP("watcher", "w", "", "Path to a Python file for event callbacks")

	carapace.Gen(launchCmd).FlagCompletion(carapace.ActionMap{
		"location":       carapace.ActionValues("default", "after", "before", "first", "hsplit", "last", "neighbor", "split", "vsplit"),
		"logo":           carapace.ActionFiles(".png"),
		"os-window-state": carapace.ActionValues("normal", "fullscreen", "maximized", "minimized"),
		"stdin-source":   carapace.ActionValues("none", "@alternate", "@alternate_scrollback", "@first_cmd_output_on_screen", "@last_cmd_output", "@last_visited_cmd_output", "@screen", "@screen_scrollback", "@selection"),
		"type":           carapace.ActionValues("window", "background", "clipboard", "os-panel", "os-window", "overlay", "overlay-main", "primary", "tab"),
		"watcher":        carapace.ActionFiles(".py"),
	})
}