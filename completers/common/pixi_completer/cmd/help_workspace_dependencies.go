package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_dependenciesCmd = &cobra.Command{
	Use:   "dependencies",
	Short: "Commands to manage workspace dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_dependenciesCmd).Standalone()
	help_workspaceCmd.AddCommand(help_workspace_dependenciesCmd)
}
