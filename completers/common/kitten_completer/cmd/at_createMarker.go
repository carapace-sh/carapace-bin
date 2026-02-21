package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createMarkerCmd = &cobra.Command{
	Use:   "create-marker",
	Short: "Create a marker that highlights specified text",
}

func init() {
	createMarkerCmd.AddCommand(atCmd)
	carapace.Gen(createMarkerCmd).Standalone()

	createMarkerCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	createMarkerCmd.Flags().StringP("match", "m", "", "The window to match")
	createMarkerCmd.Flags().Bool("self", false, "Apply marker to the window this command is run in, rather than the active window")

	carapace.Gen(createMarkerCmd).FlagCompletion(carapace.ActionMap{})
}
