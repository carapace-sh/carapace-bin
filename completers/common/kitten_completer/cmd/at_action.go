package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Run the specified mappable action",
}

func init() {
	actionCmd.AddCommand(atCmd)
	carapace.Gen(actionCmd).Standalone()

	actionCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(actionCmd).FlagCompletion(carapace.ActionMap{})
}