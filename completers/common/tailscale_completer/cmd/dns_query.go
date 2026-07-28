package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dns_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Perform a DNS query",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_queryCmd).Standalone()

	dns_queryCmd.Flags().Bool("json", false, "output in JSON format")
	dnsCmd.AddCommand(dns_queryCmd)
}
