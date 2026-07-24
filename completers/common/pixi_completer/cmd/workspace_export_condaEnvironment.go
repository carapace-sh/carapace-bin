package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_export_condaEnvironmentCmd = &cobra.Command{
	Use:   "conda-environment",
	Short: "Export workspace environment to a conda environment.yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_export_condaEnvironmentCmd).Standalone()
	workspace_export_condaEnvironmentCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	workspace_export_condaEnvironmentCmd.Flags().StringP("environment", "e", "", "The environment to render the environment file for. Defaults to the default environment")
	workspace_export_condaEnvironmentCmd.Flags().Bool("from-lock-file", false, "Render the environment with packages pinned to the versions resolved in the lock file instead of the manifest specs")
	workspace_export_condaEnvironmentCmd.Flags().StringP("name", "n", "", "The name to use for the rendered conda environment. Defaults to the environment name")
	workspace_export_condaEnvironmentCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	workspace_export_condaEnvironmentCmd.Flags().Bool("no-pypi", false, "Exclude pypi dependencies from the exported environment file")
	workspace_export_condaEnvironmentCmd.Flags().StringP("platform", "p", "", "The platform to render the environment file for. Defaults to the current platform")
	workspace_exportCmd.AddCommand(workspace_export_condaEnvironmentCmd)

	carapace.Gen(workspace_export_condaEnvironmentCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"platform":    pixi.ActionPlatforms(),
	})
}
