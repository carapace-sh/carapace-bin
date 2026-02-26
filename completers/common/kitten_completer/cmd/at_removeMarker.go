package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_removeMarkerCmd = &cobra.Command{
	Use:   "remove-marker",
	Short: "Remove the currently set marker, if any.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_removeMarkerCmd).Standalone()

	at_removeMarkerCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_removeMarkerCmd.Flags().StringP("match", "m", "", "The window to match")
	at_removeMarkerCmd.Flags().Bool("self", false, "Apply marker to the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_removeMarkerCmd)

}
