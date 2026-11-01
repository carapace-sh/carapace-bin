package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_instance_help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove stale instance entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_instance_help_pruneCmd).Standalone()

	repository_instance_helpCmd.AddCommand(repository_instance_help_pruneCmd)
}
