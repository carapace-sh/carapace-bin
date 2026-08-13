package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a domain is available to buy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_checkCmd).Standalone()

	domains_checkCmd.Flags().String("format", "", "Output format")
	domains_checkCmd.Flags().Bool("json", false, "Output as JSON")

	domainsCmd.AddCommand(domains_checkCmd)

	carapace.Gen(domains_checkCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})

	carapace.Gen(domains_checkCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
