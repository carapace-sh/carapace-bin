package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diffstat",
	Short: "make histogram from diff output",
	Long:  "https://man.freebsd.org/cgi/man.cgi?diffstat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Add SGR color")
	rootCmd.Flags().StringS("D", "D", "", "Destination directory")
	rootCmd.Flags().BoolS("E", "E", false, "Strip ANSI escapes")
	rootCmd.Flags().BoolS("K", "K", false, "Improve annotation of only files")
	rootCmd.Flags().StringS("N", "N", "", "Maximum width")
	rootCmd.Flags().BoolS("R", "R", false, "Assume swapped")
	rootCmd.Flags().StringS("S", "S", "", "Source")
	rootCmd.Flags().BoolS("T", "T", false, "Print numbers")
	rootCmd.Flags().BoolS("V", "V", false, "Version")
	rootCmd.Flags().BoolS("b", "b", false, "Ignore binary files")
	rootCmd.Flags().BoolS("c", "c", false, "Prefix with #")
	rootCmd.Flags().BoolS("d", "d", false, "Debug")
	rootCmd.Flags().StringS("e", "e", "", "Redirect stderr to file")
	rootCmd.Flags().StringS("f", "f", "", "Format")
	rootCmd.Flags().BoolS("h", "h", false, "Help")
	rootCmd.Flags().BoolS("k", "k", false, "Suppress filename merging")
	rootCmd.Flags().BoolS("l", "l", false, "List filenames only")
	rootCmd.Flags().BoolS("m", "m", false, "Merge insert/delete counts")
	rootCmd.Flags().StringS("n", "n", "", "Minimum width")
	rootCmd.Flags().StringS("o", "o", "", "Output file")
	rootCmd.Flags().StringS("p", "p", "", "Strip path prefix")
	rootCmd.Flags().BoolS("q", "q", false, "Suppress 0 files changed")
	rootCmd.Flags().StringS("r", "r", "", "Code")
	rootCmd.Flags().BoolS("s", "s", false, "Summary only")
	rootCmd.Flags().BoolS("t", "t", false, "CSV output")
	rootCmd.Flags().BoolS("u", "u", false, "Suppress filename sorting")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().StringS("w", "w", "", "Width")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionFiles(),
		"f": carapace.ActionValues("0", "1", "2", "3", "4", "5"),
		"o": carapace.ActionFiles(),
	})
}
