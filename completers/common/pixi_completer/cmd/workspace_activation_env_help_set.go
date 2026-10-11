package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_env_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_env_help_setCmd).Standalone()
	workspace_activation_env_helpCmd.AddCommand(workspace_activation_env_help_setCmd)
}
