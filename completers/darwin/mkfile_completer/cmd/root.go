package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkfile",
	Short: "create a file",
	Long:  "https://keith.github.io/xcode-manpages/mkfile.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "Create an empty filename")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
