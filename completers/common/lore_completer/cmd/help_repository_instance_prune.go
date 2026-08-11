package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_instance_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove stale instance entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_instance_pruneCmd).Standalone()

	help_repository_instanceCmd.AddCommand(help_repository_instance_pruneCmd)
}
