package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Commands to manage workspace features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_featureCmd).Standalone()

	workspaceCmd.AddCommand(workspace_featureCmd)
}
