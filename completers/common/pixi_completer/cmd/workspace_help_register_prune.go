package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_register_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prune disassociated workspaces from registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_register_pruneCmd).Standalone()

	workspace_help_registerCmd.AddCommand(workspace_help_register_pruneCmd)
}
