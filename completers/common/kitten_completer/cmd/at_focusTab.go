package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_focusTabCmd = &cobra.Command{
	Use:   "focus-tab",
	Short: "Focus the specified tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_focusTabCmd).Standalone()

	at_focusTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_focusTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_focusTabCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	atCmd.AddCommand(at_focusTabCmd)

}
