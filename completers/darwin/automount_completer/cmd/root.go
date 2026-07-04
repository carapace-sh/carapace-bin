package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "automount",
	Short: "Automount control",
	Long:  "https://keith.github.io/xcode-manpages/automount.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "Cache")
	rootCmd.Flags().StringS("t", "t", "", "Timeout")
	rootCmd.Flags().BoolS("u", "u", false, "Unmount")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
}
