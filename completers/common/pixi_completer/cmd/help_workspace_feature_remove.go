package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_feature_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a feature from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_feature_removeCmd).Standalone()

	help_workspace_featureCmd.AddCommand(help_workspace_feature_removeCmd)
}
