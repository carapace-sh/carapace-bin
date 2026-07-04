package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var mergePartitionsCmd = &cobra.Command{
	Use:   "mergePartitions",
	Short: "Merge two or more partitions into one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergePartitionsCmd).Standalone()
	rootCmd.AddCommand(mergePartitionsCmd)
	carapace.Gen(mergePartitionsCmd).PositionalCompletion(
		carapace.ActionValues("HFS+", "APFS"),
		fs.ActionBlockDevices(),
	)
	carapace.Gen(mergePartitionsCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
