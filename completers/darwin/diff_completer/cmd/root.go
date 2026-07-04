package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diff",
	Short: "differential file and directory comparator",
	Long:  "https://keith.github.io/xcode-manpages/diff.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "Ignore changes that just insert/delete blank lines")
	rootCmd.Flags().BoolS("T", "T", false, "Tab-expand")
	rootCmd.Flags().BoolS("a", "a", false, "Treat all files as text")
	rootCmd.Flags().BoolS("b", "b", false, "Ignore changes in amount of whitespace")
	rootCmd.Flags().BoolS("c", "c", false, "Context diff")
	rootCmd.Flags().BoolS("d", "d", false, "Try to find minimal set of changes")
	rootCmd.Flags().BoolS("e", "e", false, "Ed script")
	rootCmd.Flags().BoolS("f", "f", false, "Forward ed script")
	rootCmd.Flags().BoolS("i", "i", false, "Ignore case")
	rootCmd.Flags().BoolS("l", "l", false, "Long format with ASCII representation")
	rootCmd.Flags().BoolS("n", "n", false, "RCS diff")
	rootCmd.Flags().BoolS("p", "p", false, "Show C function")
	rootCmd.Flags().BoolS("q", "q", false, "Only report if files differ")
	rootCmd.Flags().BoolS("t", "t", false, "Expand tabs")
	rootCmd.Flags().BoolS("u", "u", false, "Unified diff")
	rootCmd.Flags().BoolS("w", "w", false, "Ignore all whitespace")
	rootCmd.Flags().BoolS("y", "y", false, "Side by side")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
