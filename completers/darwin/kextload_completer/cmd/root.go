package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kextload",
	Short: "Load kernel extensions",
	Long:  "https://keith.github.io/xcode-manpages/kextload.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Help|-h", "Help|-h", false, "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
