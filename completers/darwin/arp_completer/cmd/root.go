package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "arp",
	Short: "Address resolution protocol",
	Long:  "https://keith.github.io/xcode-manpages/arp.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("All|-a", "All|-a", false, "")
	rootCmd.Flags().BoolS("Delete|-d", "Delete|-d", false, "")
	rootCmd.Flags().StringS("Interface|-i", "Interface|-i", "", "")
	rootCmd.Flags().BoolS("List|-l", "List|-l", false, "")
	rootCmd.Flags().BoolS("Numeric|-n", "Numeric|-n", false, "")
	rootCmd.Flags().BoolS("Set|-s", "Set|-s", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
