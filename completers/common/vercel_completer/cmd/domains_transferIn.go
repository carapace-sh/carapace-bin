package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_transferInCmd = &cobra.Command{
	Use:   "transfer-in",
	Short: "Transfer in a domain name to Vercel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_transferInCmd).Standalone()

	domains_transferInCmd.Flags().String("code", "", "Authorization code")

	domainsCmd.AddCommand(domains_transferInCmd)

	carapace.Gen(domains_transferInCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
