package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Commands to manage workspace features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_featureCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_featureCmd)
}
