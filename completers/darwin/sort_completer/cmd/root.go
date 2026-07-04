package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "sort lines of text files",
	Long:  "https://keith.github.io/xcode-manpages/sort.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Check if input is sorted (quiet)")
	rootCmd.Flags().BoolS("M", "M", false, "Sort by month")
	rootCmd.Flags().BoolS("R", "R", false, "Random sort")
	rootCmd.Flags().BoolS("V", "V", false, "Version sort")
	rootCmd.Flags().BoolS("b", "b", false, "Ignore leading blanks")
	rootCmd.Flags().BoolS("c", "c", false, "Check if input is sorted")
	rootCmd.Flags().BoolS("d", "d", false, "Consider only blanks and alphanumeric characters")
	rootCmd.Flags().BoolS("f", "f", false, "Fold lower case to upper case")
	rootCmd.Flags().BoolS("g", "g", false, "Compare according to string numerical value")
	rootCmd.Flags().BoolS("h", "h", false, "Human numeric sort")
	rootCmd.Flags().BoolS("i", "i", false, "Consider only printable characters")
	rootCmd.Flags().BoolS("m", "m", false, "Merge already sorted files")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric sort")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse sort")
	rootCmd.Flags().BoolS("s", "s", false, "Stable sort")
	rootCmd.Flags().BoolS("u", "u", false, "Unique lines only")
	rootCmd.Flags().BoolS("z", "z", false, "Line delimiter is NUL")

	rootCmd.Flags().StringS("S", "S", "", "Set maximum memory size")
	rootCmd.Flags().StringS("T", "T", "", "Use directory for temporary files")
	rootCmd.Flags().StringS("k", "k", "", "Sort by key")
	rootCmd.Flags().StringS("o", "o", "", "Output file")
	rootCmd.Flags().StringS("t", "t", "", "Field separator")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionDirectories(),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
