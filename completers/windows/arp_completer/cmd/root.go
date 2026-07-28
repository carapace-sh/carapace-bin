package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "arp",
	Short: "display and modify the ARP cache",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/arp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "display current ARP entries")
	rootCmd.Flags().BoolP("d", "d", false, "delete an entry")
	rootCmd.Flags().BoolP("g", "g", false, "display current ARP entries (same as -a)")
	rootCmd.Flags().BoolP("s", "s", false, "add a static entry")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
