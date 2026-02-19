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

	workspace_export_condaEnvironmentCmd.Flags().StringP("environment", "e", "", "The environment to export")
	workspace_export_condaEnvironmentCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_export_condaEnvironmentCmd.Flags().StringP("name", "n", "", "The name of the conda environment")
	workspace_export_condaEnvironmentCmd.Flags().StringP("platform", "p", "", "The platform to export")
	workspace_exportCmd.AddCommand(workspace_export_condaEnvironmentCmd)

	carapace.Gen(workspace_export_condaEnvironmentCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
