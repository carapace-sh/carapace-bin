package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uniq",
	Short: "report or filter out repeated lines in a file",
	Long:  "https://keith.github.io/xcode-manpages/uniq.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Print all duplicate lines")
	rootCmd.Flags().BoolS("c", "c", false, "Precede each output line with count")
	rootCmd.Flags().BoolS("d", "d", false, "Only print duplicate lines")
	rootCmd.Flags().BoolS("i", "i", false, "Case insensitive comparison")
	rootCmd.Flags().BoolS("u", "u", false, "Only print unique lines")

	rootCmd.Flags().StringS("f", "f", "", "Skip first num fields")
	rootCmd.Flags().StringS("s", "s", "", "Skip first chars characters")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
