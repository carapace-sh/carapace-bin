package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var displaydnsCmd = &cobra.Command{
	Use:   "displaydns",
	Short: "display the DNS resolver cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displaydnsCmd).Standalone()
	rootCmd.AddCommand(displaydnsCmd)
}
