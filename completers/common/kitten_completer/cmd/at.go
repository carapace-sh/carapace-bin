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

	atCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	atCmd.Flags().String("password", "", "A password to use when contacting kitty. This will cause kitty to ask the user for permission to perform the specified action, unless the password has been accepted before or is pre-configured in kitty.conf. To use a blank password specify --use-password as always.")
	atCmd.Flags().String("password-env", "", "The name of an environment variable to read the password from. Used if no --password-file is supplied. Defaults to checking the environment variable KITTY_RC_PASSWORD.")
	atCmd.Flags().String("password-file", "", "A password to use when contacting kitty. This will cause kitty to ask the user for permission to perform the specified action, unless the password has been accepted before or is pre-configured in kitty.conf. To use a blank password specify --use-password as always.")
	atCmd.Flags().String("to", "", "An address for the kitty instance to control. Corresponds to the address given to the kitty instance via the --listen-on option or the listen_on setting in kitty.conf. If not specified, the environment variable KITTY_LISTEN_ON is checked. If that is also not found, messages are sent to the controlling terminal for this process, i.e. they will only work if this process is run within a kitty window.")
	atCmd.Flags().String("use-password", "", "If no password is available, kitty will usually just send the remote control command without a password. This option can be used to force it to always or never use the supplied password. If set to always and no password is provided, the blank password is used.")

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
