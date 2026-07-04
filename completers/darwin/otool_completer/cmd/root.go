package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otool",
	Short: "Object file display tool",
	Long:  "https://keith.github.io/xcode-manpages/otool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("All|-a", "All|-a", false, "")
	rootCmd.Flags().BoolS("Header|-h", "Header|-h", false, "")
	rootCmd.Flags().BoolS("Help|-h", "Help|-h", false, "")
	rootCmd.Flags().BoolS("Libraries|-L", "Libraries|-L", false, "")
	rootCmd.Flags().BoolS("Loadcommands|-l", "Loadcommands|-l", false, "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
