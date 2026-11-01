package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:     "bookmark snapshot|bookmark newbookmark",
	Short:   "create a bookmark of a snapshot or bookmark",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmarkCmd).Standalone()

	rootCmd.AddCommand(bookmarkCmd)

	carapace.Gen(bookmarkCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Bookmark: true}),
	)
}
