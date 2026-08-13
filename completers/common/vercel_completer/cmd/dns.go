package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Manage DNS records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dnsCmd).Standalone()

	rootCmd.AddCommand(dnsCmd)
}
