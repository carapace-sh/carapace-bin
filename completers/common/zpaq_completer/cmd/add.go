package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add archive files...",
	Short:   "append files to archive if dates have changed",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolS("force", "force", false, "append files if contents have changed")
	addCmd.Flags().BoolS("fragment", "fragment", false, "set the dedupe fragment size (64*2^N to 8128*2^N bytes)")
	addCmd.Flags().StringS("index", "index", "", "create suffix for archive indexed by F, update F")
	addCmd.Flags().StringS("key", "key", "", "create encrypted archive with password X")
	addCmd.Flags().BoolS("m0", "m0", false, "-method 0")
	addCmd.Flags().BoolS("m1", "m1", false, "-method 1")
	addCmd.Flags().BoolS("m2", "m2", false, "-method 2")
	addCmd.Flags().BoolS("m3", "m3", false, "-method 3")
	addCmd.Flags().BoolS("m4", "m4", false, "-method 4")
	addCmd.Flags().BoolS("m5", "m5", false, "-method 5")
	addCmd.Flags().StringS("method", "method", "", "compress level N (0..5 = faster..better, default: 1)")
	addCmd.Flags().StringS("summary", "summary", "", "if N > 0 show brief progress")
	rootCmd.AddCommand(addCmd)
}
