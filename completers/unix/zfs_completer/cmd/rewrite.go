package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewriteCmd = &cobra.Command{
	Use:     "rewrite [-Prvx] [-l length] [-o offset] file|directory ...",
	Short:   "rewrite specified files without modification",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewriteCmd).Standalone()

	rewriteCmd.Flags().BoolS("P", "P", false, "perform physical rewrite, preserving logical birth time")
	rewriteCmd.Flags().StringS("l", "l", "", "rewrite at most this many bytes")
	rewriteCmd.Flags().StringS("o", "o", "", "start at this offset in bytes")
	rewriteCmd.Flags().BoolS("r", "r", false, "apply recursively")
	rewriteCmd.Flags().BoolS("v", "v", false, "print names of all successfully rewritten files")
	rewriteCmd.Flags().BoolS("x", "x", false, "don't cross file system mount points when recursing")

	rootCmd.AddCommand(rewriteCmd)

	carapace.Gen(rewriteCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
