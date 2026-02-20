package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_requiresPixi_help_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_requiresPixi_help_verifyCmd).Standalone()

	workspace_requiresPixi_helpCmd.AddCommand(workspace_requiresPixi_help_verifyCmd)
}
