package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_env_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_env_setCmd).Standalone()
	help_workspace_activation_envCmd.AddCommand(help_workspace_activation_env_setCmd)
}
