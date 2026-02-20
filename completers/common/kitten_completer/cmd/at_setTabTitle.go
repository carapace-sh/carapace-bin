package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setTabTitleCmd = &cobra.Command{
	Use:   "set-tab-title",
	Short: "Set the tab title",
}

func init() {
	setTabTitleCmd.AddCommand(atCmd)
	carapace.Gen(setTabTitleCmd).Standalone()

	setTabTitleCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setTabTitleCmd.Flags().StringP("match", "m", "", "The tab to match")

	carapace.Gen(setTabTitleCmd).FlagCompletion(carapace.ActionMap{})
}