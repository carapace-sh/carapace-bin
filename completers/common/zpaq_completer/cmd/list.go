package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list archive files...",
	Short:   "list or compare external files to archive by dates",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringS("all", "all", "", "list versions in N [4] digit directories")
	listCmd.Flags().BoolS("force", "force", false, "compare file contents instead of dates")
	listCmd.Flags().StringSliceS("not", "not", nil, "exclude (* and ? match any string or char, =[+-#^?] exclude by comparison result)")
	listCmd.Flags().StringS("summary", "summary", "", "show top N sorted by size (-1: show frag IDs)")
	listCmd.Flag("all").NoOptDefVal = " "
	rootCmd.AddCommand(listCmd)
}
