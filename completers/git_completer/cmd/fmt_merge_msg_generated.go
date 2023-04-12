package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmt_merge_msgCmd = &cobra.Command{
	Use:     "fmt-merge-msg",
	Short:   "Produce a merge commit message",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(fmt_merge_msgCmd).Standalone()
	fmt_merge_msgCmd.Flags().BoolP("file", "F", false, "<file>     file to read from")
	fmt_merge_msgCmd.Flags().String("log", "", "populate log with at most <n> entries from shortlog")
	fmt_merge_msgCmd.Flags().BoolP("message", "m", false, "<text>  use <text> as start of message")
	rootCmd.AddCommand(fmt_merge_msgCmd)
}
