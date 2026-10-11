package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_script_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add activation scripts to the end of the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_script_help_addCmd).Standalone()
	workspace_activation_script_helpCmd.AddCommand(workspace_activation_script_help_addCmd)
}
