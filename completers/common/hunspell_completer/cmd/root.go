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

	rootCmd.Flags().BoolP("1", "1", false, "check only first field in lines (delimiter = tabulator)")
	rootCmd.Flags().BoolP("D", "D", false, "show available dictionaries")
	rootCmd.Flags().BoolP("G", "G", false, "print only correct words or lines")
	rootCmd.Flags().BoolP("H", "H", false, "HTML input file format")
	rootCmd.Flags().BoolP("L", "L", false, "print lines with misspelled words")
	rootCmd.Flags().BoolP("O", "O", false, "OpenDocument (ODF or Flat ODF) input file format")
	rootCmd.Flags().StringP("P", "P", "", "set password for encrypted dictionaries")
	rootCmd.Flags().BoolP("S", "S", false, "suffix words of the input text")
	rootCmd.Flags().BoolP("X", "X", false, "XML input file format")
	rootCmd.Flags().BoolP("a", "a", false, "Ispell's pipe interface")
	rootCmd.Flags().Bool("check-apostrophe", false, "check Unicode typographic apostrophe")
	rootCmd.Flags().Bool("check-url", false, "check URLs, e-mail addresses and directory paths")
	rootCmd.Flags().StringP("d", "d", "", "use d (d2 etc.) dictionaries")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("i", "i", "", "input encoding")
	rootCmd.Flags().BoolP("l", "l", false, "print misspelled words")
	rootCmd.Flags().BoolP("m", "m", false, "analyze the words of the input text")
	rootCmd.Flags().BoolP("n", "n", false, "nroff/troff input file format")
	rootCmd.Flags().StringP("p", "p", "", "set dict custom dictionary")
	rootCmd.Flags().BoolP("r", "r", false, "warn of the potential mistakes (rare words)")
	rootCmd.Flags().BoolP("s", "s", false, "stem the words of the input text")
	rootCmd.Flags().BoolP("t", "t", false, "TeX/LaTeX input file format")
	rootCmd.Flags().BoolP("version", "v", false, "print version number")
	rootCmd.Flags().BoolP("vv", "vv", false, "print Ispell compatible version number")
	rootCmd.Flags().BoolP("w", "w", false, "print misspelled words (= lines) from one word/line input")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"p": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
