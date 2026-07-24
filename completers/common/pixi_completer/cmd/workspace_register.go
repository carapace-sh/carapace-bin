package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Commands to manage the registry of workspaces. Default command will add a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_registerCmd).Standalone()

	workspace_registerCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	workspace_registerCmd.Flags().BoolP("force", "f", false, "Overwrite the workspace entry if the name of the workspace already exists in the registry")
	workspace_registerCmd.Flags().StringP("name", "n", "", "Name of the workspace to register. Defaults to the name of the current workspace")
	workspace_registerCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	workspace_registerCmd.Flags().StringP("path", "p", "", "Path to register. Defaults to the path to the current workspace")
	workspaceCmd.AddCommand(workspace_registerCmd)

	carapace.Gen(workspace_registerCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"path":        carapace.ActionDirectories(),
	})
}
