package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_env_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_env_setCmd).Standalone()
	workspace_activation_help_envCmd.AddCommand(workspace_activation_help_env_setCmd)
}
