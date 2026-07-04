package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "word, line, character, and byte count",
	Long:  "https://keith.github.io/xcode-manpages/wc.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bytes", "c", false, "Print the number of bytes in each input file")
	rootCmd.Flags().BoolP("chars", "m", false, "Print the number of characters in each input file")
	rootCmd.Flags().BoolP("lines", "l", false, "Print the number of lines in each input file")
	rootCmd.Flags().BoolP("max-line-length", "L", false, "Print the length of the longest line in each input file")
	rootCmd.Flags().BoolP("words", "w", false, "Print the number of words in each input file")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
