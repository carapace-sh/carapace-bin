package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var dumpheapCmd = &cobra.Command{
	Use:   "dumpheap",
	Short: "Dump the heap of a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpheapCmd).Standalone()

	dumpheapCmd.Flags().StringP("bitmap-format", "b", "", "dump contents of bitmaps in the given `FORMAT`")
	dumpheapCmd.Flags().BoolP("gc", "g", false, "force GC before dumping the heap")
	dumpheapCmd.Flags().BoolP("native", "n", false, "dump native heap instead of managed heap")
	dumpheapCmd.Flags().String("user", "", "specify which `USER_ID` to dump")

	rootCmd.AddCommand(dumpheapCmd)

	carapace.Gen(dumpheapCmd).FlagCompletion(carapace.ActionMap{
		"bitmap-format": carapace.ActionValues("png", "jpg", "webp"),
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(dumpheapCmd).PositionalCompletion(
		android.ActionPackages(),
		carapace.ActionFiles(),
	)
}
