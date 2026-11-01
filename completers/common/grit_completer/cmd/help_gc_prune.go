package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gc_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete unreachable loose objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gc_pruneCmd).Standalone()

	help_gcCmd.AddCommand(help_gc_pruneCmd)
}
