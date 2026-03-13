package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_getColorsCmd = &cobra.Command{
	Use:   "get-colors",
	Short: "Get terminal colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_getColorsCmd).Standalone()

	at_getColorsCmd.Flags().BoolP("configured", "c", false, "Instead of outputting the colors for the specified window, output the currently configured colors")
	at_getColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_getColorsCmd.Flags().StringP("match", "m", "", "The window to match")
	atCmd.AddCommand(at_getColorsCmd)

}
