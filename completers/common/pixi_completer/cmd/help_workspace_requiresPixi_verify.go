package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_requiresPixi_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_requiresPixi_verifyCmd).Standalone()

	help_workspace_requiresPixiCmd.AddCommand(help_workspace_requiresPixi_verifyCmd)
}
