package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Commands to manage workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_versionCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_versionCmd)
}
