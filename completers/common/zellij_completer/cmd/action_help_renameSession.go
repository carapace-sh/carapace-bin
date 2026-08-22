package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_renameSessionCmd = &cobra.Command{
	Use:   "rename-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_renameSessionCmd).Standalone()

	action_helpCmd.AddCommand(action_help_renameSessionCmd)
}
