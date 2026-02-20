package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_feature_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a feature from the manifest file",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_feature_removeCmd).Standalone()

	workspace_featureCmd.AddCommand(workspace_feature_removeCmd)
}
