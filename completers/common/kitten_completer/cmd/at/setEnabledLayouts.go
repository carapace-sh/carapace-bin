package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setEnabledLayoutsCmd = &cobra.Command{
	Use:   "set-enabled-layouts",
	Short: "Set the enabled layouts in tabs",
}

func init() {
	setEnabledLayoutsCmd.AddCommand(atCmd)
	carapace.Gen(setEnabledLayoutsCmd).Standalone()

	setEnabledLayoutsCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setEnabledLayoutsCmd).FlagCompletion(carapace.ActionMap{})
}