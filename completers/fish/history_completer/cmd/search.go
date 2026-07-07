package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search command history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolS("c", "c", false, "search by contains")
	searchCmd.Flags().Bool("case-sensitive", false, "case-sensitive search")
	searchCmd.Flags().Bool("contains", false, "search by contains")
	searchCmd.Flags().BoolP("exact", "e", false, "exact match")
	searchCmd.Flags().String("max", "", "limit results")
	searchCmd.Flags().Bool("null", false, "NUL-terminated output")
	searchCmd.Flags().BoolP("prefix", "p", false, "search by prefix")
	searchCmd.Flags().BoolP("reverse", "R", false, "oldest to newest")
	searchCmd.Flags().String("show-time", "", "prepend timestamp")
	searchCmd.Flags().BoolS("z", "z", false, "NUL-terminated output")

	rootCmd.AddCommand(searchCmd)
}
