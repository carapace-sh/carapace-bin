package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mergeIndexCmd = &cobra.Command{
	Use:     "merge-index",
	Short:   "Run a merge for files needing merging",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(mergeIndexCmd).Standalone()

	mergeIndexCmd.Flags().BoolS("a", "a", false, "run merge against all files in the index that need merging")
	mergeIndexCmd.Flags().BoolS("o", "o", false, "instead of stopping at the first failed merge, do all of them in one shot")
	mergeIndexCmd.Flags().BoolS("q", "q", false, "do not complain about a failed merge program")
	rootCmd.AddCommand(mergeIndexCmd)

	carapace.Gen(mergeIndexCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
