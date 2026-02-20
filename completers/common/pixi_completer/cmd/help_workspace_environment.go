package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Commands to manage workspace environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_environmentCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_environmentCmd)
}
