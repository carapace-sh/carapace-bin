package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "concatenate and print files",
	Long:  "https://keith.github.io/xcode-manpages/cat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "Number the non-blank lines, starting at 1")
	rootCmd.Flags().BoolS("e", "e", false, "Display non-printing characters and a $ at the end of each line")
	rootCmd.Flags().BoolS("l", "l", false, "Set an exclusive advisory lock on the standard output file descriptor")
	rootCmd.Flags().BoolS("n", "n", false, "Number the output lines, starting at 1")
	rootCmd.Flags().BoolS("s", "s", false, "Squeeze multiple adjacent empty lines into a single empty line")
	rootCmd.Flags().BoolS("t", "t", false, "Display non-printing characters and a tab at the end of each line")
	rootCmd.Flags().BoolS("u", "u", false, "Disable output buffering")
	rootCmd.Flags().BoolS("v", "v", false, "Display non-printing characters so they are visible")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
