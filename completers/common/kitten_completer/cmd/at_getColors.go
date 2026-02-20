package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getColorsCmd = &cobra.Command{
	Use:   "get-colors",
	Short: "Get terminal colors",
}

func init() {
	getColorsCmd.AddCommand(atCmd)
	carapace.Gen(getColorsCmd).Standalone()

	getColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(getColorsCmd).FlagCompletion(carapace.ActionMap{})
}