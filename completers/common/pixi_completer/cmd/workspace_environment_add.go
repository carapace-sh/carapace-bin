package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_environment_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an environment to the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_addCmd).Standalone()

	workspace_environment_addCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	workspace_environment_addCmd.Flags().Bool("force", false, "Update the environment if it already exists")
	workspace_environment_addCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_environment_addCmd.Flags().Bool("no-default-feature", false, "Don't include the default feature in the environment")
	workspace_environment_addCmd.Flags().String("solve-group", "", "The solve group for the environment")
	workspace_environmentCmd.AddCommand(workspace_environment_addCmd)

	carapace.Gen(workspace_environment_addCmd).FlagCompletion(carapace.ActionMap{
		"feature":       pixi.ActionFeatures(),
		"manifest-path": carapace.ActionFiles(),
	})
}
