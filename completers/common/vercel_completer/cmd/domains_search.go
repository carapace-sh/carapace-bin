package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Discover domain-name candidates from a keyword or fragment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_searchCmd).Standalone()

	domains_searchCmd.Flags().Bool("available", false, "Only show available domains")
	domains_searchCmd.Flags().String("format", "", "Output format")
	domains_searchCmd.Flags().Bool("json", false, "Output as JSON")
	domains_searchCmd.Flags().String("limit", "", "Number of results per page")
	domains_searchCmd.Flags().String("next", "", "Show next page of results")
	domains_searchCmd.Flags().String("order", "", "Order direction")
	domains_searchCmd.Flags().String("tld", "", "Top-level domain")

	domainsCmd.AddCommand(domains_searchCmd)

	carapace.Gen(domains_searchCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
		"order":  carapace.ActionValues("asc", "desc"),
	})
}
