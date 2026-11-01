package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the specified file in a new zellij pane with your default EDITOR Returns: Created pane ID (format: terminal_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_editCmd).Standalone()

	help_actionCmd.AddCommand(help_action_editCmd)
}
