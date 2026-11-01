package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_script_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the activation scripts in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_script_help_listCmd).Standalone()
	workspace_activation_script_helpCmd.AddCommand(workspace_activation_script_help_listCmd)
}
