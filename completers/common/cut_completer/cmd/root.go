package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cut",
	Short: "remove sections from each line of files",
	Long:  "https://man7.org/linux/man-pages/man1/cut.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bytes", "b", "", "select only these bytes")
	rootCmd.Flags().StringP("characters", "c", "", "select only these characters")
	rootCmd.Flags().Bool("complement", false, "complement the set of selected bytes, characters")
	rootCmd.Flags().StringP("delimiter", "d", "", "use DELIM instead of TAB for field delimiter")
	rootCmd.Flags().StringP("fields", "f", "", "select only these fields;  also print any line")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("n", "n", false, "(ignored)")
	rootCmd.Flags().BoolP("only-delimited", "s", false, "do not print lines not containing delimiters")
	rootCmd.Flags().String("output-delimiter", "", "use STRING as the output delimiter")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
