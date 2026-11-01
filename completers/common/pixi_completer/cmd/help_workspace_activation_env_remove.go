package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_env_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_env_removeCmd).Standalone()
	help_workspace_activation_envCmd.AddCommand(help_workspace_activation_env_removeCmd)
}
