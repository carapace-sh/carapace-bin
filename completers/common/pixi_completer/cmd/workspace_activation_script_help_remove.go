package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_script_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activation scripts from the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_script_help_removeCmd).Standalone()
	workspace_activation_script_helpCmd.AddCommand(workspace_activation_script_help_removeCmd)
}
