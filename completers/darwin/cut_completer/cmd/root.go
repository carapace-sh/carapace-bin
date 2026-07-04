package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cut",
	Short: "cut out selected portions of each line of a file",
	Long:  "https://keith.github.io/xcode-manpages/cut.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "Do not split multi-byte characters")
	rootCmd.Flags().BoolS("s", "s", false, "Do not print lines not containing delimiters")
	rootCmd.Flags().BoolS("w", "w", false, "Use whitespace (spaces and tabs) as the delimiter")

	rootCmd.Flags().StringS("b", "b", "", "Select only these bytes")
	rootCmd.Flags().StringS("c", "c", "", "Select only these characters")
	rootCmd.Flags().StringS("d", "d", "", "Use DELIM instead of TAB for field delimiter")
	rootCmd.Flags().StringS("f", "f", "", "Select only these fields")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
