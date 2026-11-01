package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Commands to manage the activation scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_scriptCmd).Standalone()
	help_workspace_activationCmd.AddCommand(help_workspace_activation_scriptCmd)
}
