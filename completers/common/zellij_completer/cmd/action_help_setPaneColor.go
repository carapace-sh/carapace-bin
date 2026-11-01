package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_setPaneColorCmd = &cobra.Command{
	Use:   "set-pane-color",
	Short: "Set the default foreground/background color of a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_setPaneColorCmd).Standalone()

	action_helpCmd.AddCommand(action_help_setPaneColorCmd)
}
