package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_renameSessionCmd = &cobra.Command{
	Use:   "rename-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_renameSessionCmd).Standalone()

	help_actionCmd.AddCommand(help_action_renameSessionCmd)
}
