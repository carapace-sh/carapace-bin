package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_createMarkerCmd = &cobra.Command{
	Use:   "create-marker",
	Short: "Create a marker that highlights specified text",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_createMarkerCmd).Standalone()

	at_createMarkerCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_createMarkerCmd.Flags().StringP("match", "m", "", "The window to match")
	at_createMarkerCmd.Flags().Bool("self", false, "Apply marker to the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_createMarkerCmd)

}
