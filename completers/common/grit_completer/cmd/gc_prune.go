package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete unreachable loose objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_pruneCmd).Standalone()

	gc_pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gc_pruneCmd.Flags().String("older-than", "", "Only remove objects older than this date (Git approxidate syntax)")
	gcCmd.AddCommand(gc_pruneCmd)
}
