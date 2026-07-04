package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tr",
	Short: "translate characters",
	Long:  "https://keith.github.io/xcode-manpages/tr.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Complement the set of characters in string1")
	rootCmd.Flags().BoolS("c", "c", false, "Complement the set of characters in string1")
	rootCmd.Flags().BoolS("d", "d", false, "Delete characters in string1")
	rootCmd.Flags().BoolS("s", "s", false, "Squeeze repeats")
	rootCmd.Flags().BoolS("u", "u", false, "Handle binary data")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
