package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_register_pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Prune disassociated workspaces from registry",
	Aliases: []string{"pr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_register_pruneCmd).Standalone()

	workspace_registerCmd.AddCommand(workspace_register_pruneCmd)
}
