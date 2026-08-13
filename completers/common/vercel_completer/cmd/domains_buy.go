package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Purchase a new domain name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_buyCmd).Standalone()

	domainsCmd.AddCommand(domains_buyCmd)

	carapace.Gen(domains_buyCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
