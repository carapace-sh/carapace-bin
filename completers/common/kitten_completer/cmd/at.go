package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var atCmd = &cobra.Command{
	Use:   "@",
	Short: "Control kitty remotely",
}

func init() {
	rootCmd.AddCommand(atCmd)
	carapace.Gen(atCmd).Standalone()

	atCmd.Flags().String("to", "", "An address for the kitty instance to control. Corresponds to the address given to the kitty instance via the --listen-on option or the listen_on setting in kitty.conf. If not specified, the environment variable KITTY_LISTEN_ON is checked. If that is also not found, messages are sent to the controlling terminal for this process, i.e. they will only work if this process is run within a kitty window.")
	atCmd.Flags().String("password", "", "A password to use when contacting kitty. This will cause kitty to ask the user for permission to perform the specified action, unless the password has been accepted before or is pre-configured in kitty.conf. To use a blank password specify --use-password as always.")
	atCmd.Flags().String("password-file", "", "A password to use when contacting kitty. This will cause kitty to ask the user for permission to perform the specified action, unless the password has been accepted before or is pre-configured in kitty.conf. To use a blank password specify --use-password as always.")
	atCmd.Flags().String("password-env", "", "The name of an environment variable to read the password from. Used if no --password-file is supplied. Defaults to checking the environment variable KITTY_RC_PASSWORD.")
	atCmd.Flags().String("use-password", "", "If no password is available, kitty will usually just send the remote control command without a password. This option can be used to force it to always or never use the supplied password. If set to always and no password is provided, the blank password is used.")
	atCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	atCmd.MarkFlagsMutuallyExclusive("password-file", "password-env")

	carapace.Gen(atCmd).FlagCompletion(carapace.ActionMap{
		"password-file": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed("-", "Read from STDIN"),
			carapace.ActionValuesDescribed("fd:", "Read from a file descriptor").NoSpace(':'),
		).ToA(),
		"password-env": os.ActionEnvironmentVariables(),
		"use-password": carapace.ActionValues("if-available", "always", "never"),
	})
}

// TODO: subcommands
// action:Run the specified mappable action
// close-tab:Close the specified tabs
// close-window:Close the specified windows
// create-marker:Create a marker that highlights specified text
// detach-tab:Detach the specified tabs and place them in a different/new OS window
// detach-window:Detach the specified windows and place them in a different/new tab
// disable-ligatures:Control ligature rendering
// env:Change environment variables seen by future children
// focus-tab:Focus the specified tab
// focus-window:Focus the specified window
// get-colors:Get terminal colors
// get-text:Get text from the specified window
// goto-layout:Set the window layout
// kitten:Run a kitten
// last-used-layout:Switch to the last used layout
// launch:Run an arbitrary process in a new window/tab
// load-config:(Re)load a config file
// ls:List tabs/windows
// new-window:Open new window
// remove-marker:Remove the currently set marker, if any.
// resize-os-window:Resize/show/hide/etc. the specified OS Windows
// resize-window:Resize the specified windows
// run:Run a program on the computer in which kitty is running and get the output
// scroll-window:Scroll the specified windows
// select-window:Visually select a window in the specified tab
// send-key:Send arbitrary key presses to the specified windows
// send-text:Send arbitrary text to specified windows
// set-background-image:Set the background image
// set-background-opacity:Set the background opacity
// set-colors:Set terminal colors
// set-enabled-layouts:Set the enabled layouts in tabs
// set-font-size:Set the font size in the active top-level OS window
// set-spacing:Set window paddings and margins
// set-tab-color:Change the color of the specified tabs in the tab bar
// set-tab-title:Set the tab title
// set-user-vars:Set user variables on a window
// set-window-logo:Set the window logo
// set-window-title:Set the window title
// signal-child:Send a signal to the foreground process in the specified windows
