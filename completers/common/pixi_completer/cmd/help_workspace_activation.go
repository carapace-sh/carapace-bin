package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activationCmd = &cobra.Command{
	Use:   "activation",
	Short: "Commands to manage environment activation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activationCmd).Standalone()
	help_workspaceCmd.AddCommand(help_workspace_activationCmd)
}
