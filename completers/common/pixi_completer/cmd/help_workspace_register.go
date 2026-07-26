package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Commands to manage the registry of workspaces. Default command will add a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_registerCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_registerCmd)
}
