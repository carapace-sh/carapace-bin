package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_focusWindowCmd = &cobra.Command{
	Use:   "focus-window",
	Short: "Focus the specified window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_focusWindowCmd).Standalone()

	at_focusWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_focusWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_focusWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")
	atCmd.AddCommand(at_focusWindowCmd)

}
