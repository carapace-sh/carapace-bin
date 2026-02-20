package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setSpacingCmd = &cobra.Command{
	Use:   "set-spacing",
	Short: "Set window paddings and margins",
}

func init() {
	setSpacingCmd.AddCommand(atCmd)
	carapace.Gen(setSpacingCmd).Standalone()

	setSpacingCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setSpacingCmd).FlagCompletion(carapace.ActionMap{})
}