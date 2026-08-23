package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hiutil",
	Short: "help indexer utility",
	Long:  "https://man.freebsd.org/cgi/man.cgi?hiutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "Index one file at a time")
	rootCmd.Flags().BoolS("A", "A", false, "List anchors")
	rootCmd.Flags().BoolS("C", "C", false, "Create")
	rootCmd.Flags().BoolS("D", "D", false, "List anchor dictionary")
	rootCmd.Flags().BoolS("E", "E", false, "List index versions")
	rootCmd.Flags().BoolS("F", "F", false, "List files")
	rootCmd.Flags().BoolS("H", "H", false, "Help")
	rootCmd.Flags().BoolS("I", "I", false, "Index format")
	rootCmd.Flags().BoolS("M", "M", false, "List min term length")
	rootCmd.Flags().BoolS("P", "P", false, "Purge caches")
	rootCmd.Flags().BoolS("S", "S", false, "List stopwords")
	rootCmd.Flags().BoolS("T", "T", false, "List terms")
	rootCmd.Flags().BoolS("V", "V", false, "Version")
	rootCmd.Flags().BoolS("a", "a", false, "Anchors")
	rootCmd.Flags().StringS("e", "e", "", "Exclude pattern")
	rootCmd.Flags().StringS("f", "f", "", "File")
	rootCmd.Flags().BoolS("g", "g", false, "Generate summaries")
	rootCmd.Flags().StringS("i", "i", "", "Include pattern")
	rootCmd.Flags().StringS("l", "l", "", "Locale")
	rootCmd.Flags().StringS("m", "m", "", "Min term length")
	rootCmd.Flags().StringS("r", "r", "", "Remote URL")
	rootCmd.Flags().StringS("s", "s", "", "Stopwords")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
}
