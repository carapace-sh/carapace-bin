package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_feature_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the features in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_feature_listCmd).Standalone()

	help_workspace_featureCmd.AddCommand(help_workspace_feature_listCmd)
}
