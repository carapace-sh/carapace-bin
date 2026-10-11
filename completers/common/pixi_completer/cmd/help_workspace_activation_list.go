package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List activation scripts and environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_listCmd).Standalone()
	help_workspace_activationCmd.AddCommand(help_workspace_activation_listCmd)
}
