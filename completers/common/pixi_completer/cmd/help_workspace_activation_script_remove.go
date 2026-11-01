package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_script_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activation scripts from the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_script_removeCmd).Standalone()
	help_workspace_activation_scriptCmd.AddCommand(help_workspace_activation_script_removeCmd)
}
