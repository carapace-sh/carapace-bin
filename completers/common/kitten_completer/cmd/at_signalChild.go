package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signalChildCmd = &cobra.Command{
	Use:   "signal-child",
	Short: "Send a signal to the foreground process in the specified windows",
}

func init() {
	signalChildCmd.AddCommand(atCmd)
	carapace.Gen(signalChildCmd).Standalone()

	signalChildCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	signalChildCmd.Flags().StringP("match", "m", "", "The window to match")
	signalChildCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")

	carapace.Gen(signalChildCmd).FlagCompletion(carapace.ActionMap{})
}
