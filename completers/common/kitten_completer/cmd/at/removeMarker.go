package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeMarkerCmd = &cobra.Command{
	Use:   "remove-marker",
	Short: "Remove the currently set marker, if any.",
}

func init() {
	removeMarkerCmd.AddCommand(atCmd)
	carapace.Gen(removeMarkerCmd).Standalone()

	removeMarkerCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(removeMarkerCmd).FlagCompletion(carapace.ActionMap{})
}