package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_requiresPixi_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_requiresPixi_unsetCmd).Standalone()

	workspace_help_requiresPixiCmd.AddCommand(workspace_help_requiresPixi_unsetCmd)
}
