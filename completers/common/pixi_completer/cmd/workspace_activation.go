package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activationCmd = &cobra.Command{
	Use:   "activation",
	Short: "Commands to manage environment activation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activationCmd).Standalone()

	workspace_activationCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths")
	workspace_activationCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files")
	workspaceCmd.AddCommand(workspace_activationCmd)

	carapace.Gen(workspace_activationCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
