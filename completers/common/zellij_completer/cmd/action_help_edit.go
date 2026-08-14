package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the specified file in a new zellij pane with your default EDITOR Returns: Created pane ID (format: terminal_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_editCmd).Standalone()

	action_helpCmd.AddCommand(action_help_editCmd)
}
