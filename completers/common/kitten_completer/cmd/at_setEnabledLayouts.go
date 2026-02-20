package cmd

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

	setEnabledLayoutsCmd.Flags().Bool("configured", false, "Change the default enabled layout value so that the new value takes effect for all newly created tabs as well")
	setEnabledLayoutsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setEnabledLayoutsCmd.Flags().StringP("match", "m", "", "The tab to match")

	carapace.Gen(setEnabledLayoutsCmd).FlagCompletion(carapace.ActionMap{})
}