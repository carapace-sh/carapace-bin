package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_env_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the activation environment variables in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_env_listCmd).Standalone()
	workspace_activation_help_envCmd.AddCommand(workspace_activation_help_env_listCmd)
}
