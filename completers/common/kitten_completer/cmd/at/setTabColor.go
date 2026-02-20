package at

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

	carapace.Gen(setTabColorCmd).FlagCompletion(carapace.ActionMap{})
}