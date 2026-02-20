package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gotoLayoutCmd = &cobra.Command{
	Use:   "goto-layout",
	Short: "Set the window layout",
}

func init() {
	gotoLayoutCmd.AddCommand(atCmd)
	carapace.Gen(gotoLayoutCmd).Standalone()

	gotoLayoutCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(gotoLayoutCmd).FlagCompletion(carapace.ActionMap{})
}