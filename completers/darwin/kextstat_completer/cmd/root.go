package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kextstat",
	Short: "List loaded kernel extensions",
	Long:  "https://keith.github.io/xcode-manpages/kextstat.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Show all kexts")
	rootCmd.Flags().StringS("b", "b", "", "Bundle identifier")
	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().BoolS("k", "k", false, "Show kexts only")
	rootCmd.Flags().BoolS("l", "l", false, "Long format")
}
