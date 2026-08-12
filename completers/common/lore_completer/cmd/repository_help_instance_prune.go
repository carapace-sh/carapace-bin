package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_instance_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove stale instance entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_instance_pruneCmd).Standalone()

	repository_help_instanceCmd.AddCommand(repository_help_instance_pruneCmd)
}
