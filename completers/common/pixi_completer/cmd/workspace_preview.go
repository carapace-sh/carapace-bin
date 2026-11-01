package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Commands to manage workspace preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_previewCmd).Standalone()

	workspace_previewCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths")
	workspace_previewCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files")
	workspaceCmd.AddCommand(workspace_previewCmd)

	carapace.Gen(workspace_previewCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
