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

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"a", "display current ARP entries",
			"g", "display current ARP entries (same as -a)",
			"d", "delete an entry",
			"s", "add a static entry",
		),
	)
}
