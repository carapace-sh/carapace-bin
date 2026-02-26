package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a program on the computer in which kitty is running and get the output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_runCmd).Standalone()

	at_runCmd.Flags().Bool("allow-remote-control", false, "The executed program will have privileges to run remote control commands in kitty")
	at_runCmd.Flags().String("env", "", "Environment variables to set in the child process")
	at_runCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_runCmd.Flags().String("remote-control-password", "", "Restrict the actions remote control is allowed to take")
	atCmd.AddCommand(at_runCmd)

}
