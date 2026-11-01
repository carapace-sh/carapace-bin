package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_editScrollbackCmd = &cobra.Command{
	Use:   "edit-scrollback",
	Short: "Open the pane scrollback in your default editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_editScrollbackCmd).Standalone()

	help_actionCmd.AddCommand(help_action_editScrollbackCmd)
}
