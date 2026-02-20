package at

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

	carapace.Gen(signalChildCmd).FlagCompletion(carapace.ActionMap{})
}