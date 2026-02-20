package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Commands to manage workspace features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_featureCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_featureCmd)
}
