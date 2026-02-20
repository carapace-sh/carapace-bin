package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_feature_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a feature from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_feature_help_removeCmd).Standalone()

	workspace_feature_helpCmd.AddCommand(workspace_feature_help_removeCmd)
}
