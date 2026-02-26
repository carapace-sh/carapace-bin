package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_scrollWindowCmd = &cobra.Command{
	Use:   "scroll-window",
	Short: "Scroll the specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_scrollWindowCmd).Standalone()

	at_scrollWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_scrollWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_scrollWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	atCmd.AddCommand(at_scrollWindowCmd)

}
