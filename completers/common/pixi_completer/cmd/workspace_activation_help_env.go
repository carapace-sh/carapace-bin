package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Commands to manage activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_envCmd).Standalone()
	workspace_activation_helpCmd.AddCommand(workspace_activation_help_envCmd)
}
