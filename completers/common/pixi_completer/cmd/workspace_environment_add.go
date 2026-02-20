package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_environment_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds an environment to the manifest file",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_addCmd).Standalone()

	workspace_environment_addCmd.Flags().StringSliceP("feature", "f", nil, "Features to add to the environment")
	workspace_environment_addCmd.Flags().Bool("force", false, "Update the manifest even if the environment already exists")
	workspace_environment_addCmd.Flags().Bool("no-default-feature", false, "Don't include the default feature in the environment")
	workspace_environment_addCmd.Flags().String("solve-group", "", "The solve-group to add the environment to")
	workspace_environmentCmd.AddCommand(workspace_environment_addCmd)

	carapace.Gen(workspace_environment_addCmd).FlagCompletion(carapace.ActionMap{
		"feature": pixi.ActionFeatures(),
	})
}
