package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_script_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add activation scripts to the end of the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_script_addCmd).Standalone()
	workspace_help_activation_scriptCmd.AddCommand(workspace_help_activation_script_addCmd)
}
