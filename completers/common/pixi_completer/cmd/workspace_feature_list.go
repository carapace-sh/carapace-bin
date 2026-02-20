package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_feature_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the features in the manifest file",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_feature_listCmd).Standalone()

	workspace_featureCmd.AddCommand(workspace_feature_listCmd)
}
