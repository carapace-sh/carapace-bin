package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_script_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activation scripts from the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_script_removeCmd).Standalone()
	workspace_activation_help_scriptCmd.AddCommand(workspace_activation_help_script_removeCmd)
}
