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

	workspace_export_condaEnvironmentCmd.Flags().StringP("environment", "e", "", "The environment to render the environment file for. Defaults to the default environment")
	workspace_export_condaEnvironmentCmd.Flags().StringP("name", "n", "", "The name to use for the rendered conda environment. Defaults to the environment name")
	workspace_export_condaEnvironmentCmd.Flags().StringP("platform", "p", "", "The platform to render the environment file for. Defaults to the current platform")
	workspace_exportCmd.AddCommand(workspace_export_condaEnvironmentCmd)

	carapace.Gen(workspace_export_condaEnvironmentCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"platform":    pixi.ActionPlatforms(),
	})
}
