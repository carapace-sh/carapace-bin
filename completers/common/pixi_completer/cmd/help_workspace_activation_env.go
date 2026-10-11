package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_activation_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Commands to manage activation environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_activation_envCmd).Standalone()
	help_workspace_activationCmd.AddCommand(help_workspace_activation_envCmd)
}
