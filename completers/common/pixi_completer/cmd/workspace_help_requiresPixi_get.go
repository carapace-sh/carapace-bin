package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_requiresPixi_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_requiresPixi_getCmd).Standalone()

	workspace_help_requiresPixiCmd.AddCommand(workspace_help_requiresPixi_getCmd)
}
