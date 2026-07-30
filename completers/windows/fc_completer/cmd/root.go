package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc",
	Short: "compare two files and display the differences",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "ASCII text comparison")
	rootCmd.Flags().BoolP("b", "b", false, "binary comparison")
	rootCmd.Flags().BoolP("c", "c", false, "case-insensitive comparison")
	rootCmd.Flags().BoolP("l", "l", false, "compare files as ASCII text")
	rootCmd.Flags().BoolP("n", "n", false, "display line numbers")
	rootCmd.Flags().BoolP("t", "t", false, "do not expand tabs to spaces")
	rootCmd.Flags().BoolP("u", "u", false, "compare files as Unicode text")
	rootCmd.Flags().BoolP("w", "w", false, "compress whitespace (tabs and spaces)")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
