package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_dependenciesCmd = &cobra.Command{
	Use:   "dependencies",
	Short: "Commands to manage workspace dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependenciesCmd).Standalone()

	workspace_dependenciesCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths")
	workspace_dependenciesCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files")
	workspaceCmd.AddCommand(workspace_dependenciesCmd)

	carapace.Gen(workspace_dependenciesCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
