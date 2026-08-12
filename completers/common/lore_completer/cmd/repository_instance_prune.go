package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_instance_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove stale instance entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_instance_pruneCmd).Standalone()

	repository_instance_pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_instanceCmd.AddCommand(repository_instance_pruneCmd)
}
