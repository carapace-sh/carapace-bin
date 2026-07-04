package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tee",
	Short: "read from standard input and write to standard output and files",
	Long:  "https://keith.github.io/xcode-manpages/tee.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Append to output file")
	rootCmd.Flags().BoolS("i", "i", false, "Ignore interrupt signals")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
