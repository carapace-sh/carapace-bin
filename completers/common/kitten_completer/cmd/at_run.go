package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a program on the computer in which kitty is running and get the output",
}

func init() {
	runCmd.AddCommand(atCmd)
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	runCmd.Flags().Bool("allow-remote-control", false, "The executed program will have privileges to run remote control commands in kitty")
	runCmd.Flags().String("env", "", "Environment variables to set in the child process")
	runCmd.Flags().String("remote-control-password", "", "Restrict the actions remote control is allowed to take")

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{})
}