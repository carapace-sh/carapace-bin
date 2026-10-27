package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete unreachable loose objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_help_pruneCmd).Standalone()

	gc_helpCmd.AddCommand(gc_help_pruneCmd)
}
