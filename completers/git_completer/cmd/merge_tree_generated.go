package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var merge_treeCmd = &cobra.Command{
	Use:     "merge-tree",
	Short:   "Show three-way merge without touching index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(merge_treeCmd).Standalone()

	rootCmd.AddCommand(merge_treeCmd)
}
