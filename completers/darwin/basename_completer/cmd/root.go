package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "basename",
	Short: "return filename or directory portion of pathname",
	Long:  "https://keith.github.io/xcode-manpages/basename.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Treat each argument as a separate string")

	rootCmd.Flags().StringS("s", "s", "", "Remove trailing suffix")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
