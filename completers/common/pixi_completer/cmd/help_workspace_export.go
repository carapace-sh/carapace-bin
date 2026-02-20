package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Commands to export workspaces to other formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_exportCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_exportCmd)
}
