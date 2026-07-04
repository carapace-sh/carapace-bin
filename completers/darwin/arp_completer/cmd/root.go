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

	rootCmd.Flags().BoolS("a", "a", false, "Display all ARP entries")
	rootCmd.Flags().BoolS("d", "d", false, "Delete an ARP entry")
	rootCmd.Flags().BoolS("l", "l", false, "Show link-layer address")
	rootCmd.Flags().BoolS("n", "n", false, "Do not resolve host names")
	rootCmd.Flags().BoolS("s", "s", false, "Set an ARP entry")

	rootCmd.Flags().StringS("i", "i", "", "Specify interface")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
