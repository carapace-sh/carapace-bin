package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dns_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List DNS entries",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_lsCmd).Standalone()

	dns_lsCmd.Flags().String("limit", "", "Number of results per page")
	dns_lsCmd.Flags().String("next", "", "Show next page of results")

	dnsCmd.AddCommand(dns_lsCmd)

	carapace.Gen(dns_lsCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
