package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List activation scripts and environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_listCmd).Standalone()
	workspace_activation_helpCmd.AddCommand(workspace_activation_help_listCmd)
}
