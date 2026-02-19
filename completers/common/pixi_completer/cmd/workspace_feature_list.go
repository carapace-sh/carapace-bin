package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_feature_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the features in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_feature_listCmd).Standalone()

	workspace_feature_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_featureCmd.AddCommand(workspace_feature_listCmd)

	carapace.Gen(workspace_feature_listCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
