package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_script_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add activation scripts to the end of the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_script_addCmd).Standalone()
	help_workspace_activation_scriptCmd.AddCommand(help_workspace_activation_script_addCmd)
}
