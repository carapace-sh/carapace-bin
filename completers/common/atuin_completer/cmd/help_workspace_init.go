package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a .atuin.lua workspace config file in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_initCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_initCmd)
}
