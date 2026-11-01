package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List activation scripts and environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_listCmd).Standalone()
	workspace_help_activationCmd.AddCommand(workspace_help_activation_listCmd)
}
