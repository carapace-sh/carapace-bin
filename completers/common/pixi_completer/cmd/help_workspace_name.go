package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Commands to manage workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_nameCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_nameCmd)
}
