package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agvtool",
	Short: "Apple Generic Versioning tool",
	Long:  "https://keith.github.io/xcode-manpages/agvtool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Help|-h", "Help|-h", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
