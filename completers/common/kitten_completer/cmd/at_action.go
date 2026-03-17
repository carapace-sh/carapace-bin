package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kitty"
	"github.com/spf13/cobra"
)

var at_actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Run the specified mappable action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_actionCmd).Standalone()

	at_actionCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_actionCmd.Flags().StringP("match", "m", "", "The window to match")
	at_actionCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	at_actionCmd.Flags().Bool("self", false, "Run the action on the window this command is run in instead of the active window")
	atCmd.AddCommand(at_actionCmd)

	carapace.Gen(at_actionCmd).PositionalAnyCompletion(
		kitty.ActionActions(),
	)
}
