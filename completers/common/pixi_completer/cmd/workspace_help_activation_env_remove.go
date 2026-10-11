package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_env_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_env_removeCmd).Standalone()
	workspace_help_activation_envCmd.AddCommand(workspace_help_activation_env_removeCmd)
}
