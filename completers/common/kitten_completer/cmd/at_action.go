package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kitty"
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
	actionCmd.Flags().StringP("match", "m", "", "The window to match")
	actionCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	actionCmd.Flags().Bool("self", false, "Run the action on the window this command is run in instead of the active window")

	carapace.Gen(actionCmd).FlagCompletion(carapace.ActionMap{})

	carapace.Gen(actionCmd).PositionalAnyCompletion(
		kitty.ActionActions(),
	)
}
