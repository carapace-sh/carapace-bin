package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Commands to manage the activation scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_scriptCmd).Standalone()
	workspace_help_activationCmd.AddCommand(workspace_help_activation_scriptCmd)
}
