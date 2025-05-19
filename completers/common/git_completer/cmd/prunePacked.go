package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prunePackedCmd = &cobra.Command{
	Use:     "prune-packed",
	Short:   "Remove extra objects that are already in pack files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(prunePackedCmd).Standalone()

	prunePackedCmd.Flags().BoolP("dry-run", "n", false, "only show objects that would have been removed")
	prunePackedCmd.Flags().BoolP("quiet", "q", false, "squelch the progress indicator")
	rootCmd.AddCommand(prunePackedCmd)
}
