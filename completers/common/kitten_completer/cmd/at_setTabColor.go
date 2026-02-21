package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setTabColorCmd = &cobra.Command{
	Use:   "set-tab-color",
	Short: "Change the color of the specified tabs in the tab bar",
}

func init() {
	setTabColorCmd.AddCommand(atCmd)
	carapace.Gen(setTabColorCmd).Standalone()

	setTabColorCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setTabColorCmd.Flags().StringP("match", "m", "", "The tab to match")
	setTabColorCmd.Flags().Bool("self", false, "Close the tab this command is run in, rather than the active tab")

	carapace.Gen(setTabColorCmd).FlagCompletion(carapace.ActionMap{})
}
