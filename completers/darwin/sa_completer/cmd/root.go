package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sa",
	Short: "print system accounting statistics",
	Long:  "https://keith.github.io/xcode-manpages/sa.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Sort by total number of disk I/O operations")
	rootCmd.Flags().BoolS("K", "K", false, "Sort by CPU-storage integral")
	rootCmd.Flags().StringS("P", "P", "", "Print only commands named in file")
	rootCmd.Flags().StringS("U", "U", "", "Print only commands owned by users in file")
	rootCmd.Flags().BoolS("a", "a", false, "List all command names")
	rootCmd.Flags().BoolS("b", "b", false, "Sort by user and system time")
	rootCmd.Flags().BoolS("c", "c", false, "Sort by total CPU time")
	rootCmd.Flags().BoolS("d", "d", false, "Sort by average number of disk I/O operations")
	rootCmd.Flags().BoolS("f", "f", false, "Force no interactive threshold")
	rootCmd.Flags().BoolS("i", "i", false, "Do not read summary files")
	rootCmd.Flags().BoolS("j", "j", false, "Sort by CPU time per execution")
	rootCmd.Flags().BoolS("k", "k", false, "Sort by CPU-time averaged core usage")
	rootCmd.Flags().BoolS("l", "l", false, "Separate system and user time")
	rootCmd.Flags().BoolS("m", "m", false, "Sort by number of processes")
	rootCmd.Flags().BoolS("n", "n", false, "Sort by number of invocations")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse sort order")
	rootCmd.Flags().BoolS("s", "s", false, "Merge accounting files")
	rootCmd.Flags().BoolS("t", "t", false, "Sort by real time")
	rootCmd.Flags().BoolS("u", "u", false, "Suppress printing of command names")
	rootCmd.Flags().StringS("v", "v", "", "Cutoff value")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"P": carapace.ActionFiles(),
		"U": carapace.ActionFiles(),
	})
}
