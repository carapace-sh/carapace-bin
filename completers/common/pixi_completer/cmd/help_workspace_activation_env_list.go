package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_env_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the activation environment variables in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_env_listCmd).Standalone()
	help_workspace_activation_envCmd.AddCommand(help_workspace_activation_env_listCmd)
}
