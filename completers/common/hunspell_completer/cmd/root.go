package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hunspell",
	Short: "spell checker",
	Long:  "https://hunspell.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "check only first field in lines (delimiter = tabulator)")
	rootCmd.Flags().BoolS("D", "D", false, "show available dictionaries")
	rootCmd.Flags().BoolS("G", "G", false, "print only correct words or lines")
	rootCmd.Flags().BoolS("H", "H", false, "HTML input file format")
	rootCmd.Flags().BoolS("L", "L", false, "print lines with misspelled words")
	rootCmd.Flags().BoolS("O", "O", false, "OpenDocument (ODF or Flat ODF) input file format")
	rootCmd.Flags().StringS("P", "P", "", "set password for encrypted dictionaries")
	rootCmd.Flags().BoolS("S", "S", false, "suffix words of the input text")
	rootCmd.Flags().BoolS("X", "X", false, "XML input file format")
	rootCmd.Flags().BoolS("a", "a", false, "Ispell's pipe interface")
	rootCmd.Flags().Bool("check-apostrophe", false, "check Unicode typographic apostrophe")
	rootCmd.Flags().Bool("check-url", false, "check URLs, e-mail addresses and directory paths")
	rootCmd.Flags().StringS("d", "d", "", "use d (d2 etc.) dictionaries")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringS("i", "i", "", "input encoding")
	rootCmd.Flags().BoolS("l", "l", false, "print misspelled words")
	rootCmd.Flags().BoolS("m", "m", false, "analyze the words of the input text")
	rootCmd.Flags().BoolS("n", "n", false, "nroff/troff input file format")
	rootCmd.Flags().StringS("p", "p", "", "set dict custom dictionary")
	rootCmd.Flags().BoolS("r", "r", false, "warn of the potential mistakes (rare words)")
	rootCmd.Flags().BoolS("s", "s", false, "stem the words of the input text")
	rootCmd.Flags().BoolS("t", "t", false, "TeX/LaTeX input file format")
	rootCmd.Flags().BoolP("version", "v", false, "print version number")
	rootCmd.Flags().BoolS("vv", "vv", false, "print Ispell compatible version number")
	rootCmd.Flags().BoolS("w", "w", false, "print misspelled words (= lines) from one word/line input")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"p": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
