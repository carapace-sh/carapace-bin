package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dns_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove a DNS entry using its ID",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_rmCmd).Standalone()

	dns_rmCmd.Flags().Bool("yes", false, "Skip confirmation")

	dnsCmd.AddCommand(dns_rmCmd)

	carapace.Gen(dns_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
