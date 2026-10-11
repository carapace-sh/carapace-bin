package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_script_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the activation scripts in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_script_listCmd).Standalone()
	workspace_help_activation_scriptCmd.AddCommand(workspace_help_activation_script_listCmd)
}
