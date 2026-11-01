package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_editScrollbackCmd = &cobra.Command{
	Use:   "edit-scrollback",
	Short: "Open the pane scrollback in your default editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_editScrollbackCmd).Standalone()

	action_helpCmd.AddCommand(action_help_editScrollbackCmd)
}
