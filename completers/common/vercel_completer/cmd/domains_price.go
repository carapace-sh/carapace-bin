package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Show registrar price quotes for one or more domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_priceCmd).Standalone()

	domains_priceCmd.Flags().String("format", "", "Output format")
	domains_priceCmd.Flags().Bool("json", false, "Output as JSON")

	domainsCmd.AddCommand(domains_priceCmd)

	carapace.Gen(domains_priceCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
