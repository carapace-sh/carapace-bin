package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_feature_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a feature from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_feature_removeCmd).Standalone()

	workspace_feature_removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_featureCmd.AddCommand(workspace_feature_removeCmd)

	carapace.Gen(workspace_feature_removeCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
