package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Commands to export workspaces to other formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_exportCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_exportCmd)
}
