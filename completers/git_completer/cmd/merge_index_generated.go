package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var merge_indexCmd = &cobra.Command{
	Use:     "merge-index",
	Short:   "Run a merge for files needing merging",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(merge_indexCmd).Standalone()

	rootCmd.AddCommand(merge_indexCmd)
}
