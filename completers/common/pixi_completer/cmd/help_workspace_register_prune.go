package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_register_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prune disassociated workspaces from registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_register_pruneCmd).Standalone()

	help_workspace_registerCmd.AddCommand(help_workspace_register_pruneCmd)
}
