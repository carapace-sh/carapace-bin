package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_signalChildCmd = &cobra.Command{
	Use:   "signal-child",
	Short: "Send a signal to the foreground process in the specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_signalChildCmd).Standalone()

	at_signalChildCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_signalChildCmd.Flags().StringP("match", "m", "", "The window to match")
	at_signalChildCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	atCmd.AddCommand(at_signalChildCmd)

}
