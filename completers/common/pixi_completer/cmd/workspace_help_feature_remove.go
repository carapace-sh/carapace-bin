package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_feature_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a feature from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_feature_removeCmd).Standalone()

	workspace_help_featureCmd.AddCommand(workspace_help_feature_removeCmd)
}
