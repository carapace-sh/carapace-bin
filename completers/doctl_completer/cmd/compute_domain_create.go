package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a domain to your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_createCmd).Standalone()
	compute_domain_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Domain`, `TTL`")
	compute_domain_createCmd.Flags().String("ip-address", "", "Creates an A record when an IPv4 address is provided")
	compute_domain_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_domainCmd.AddCommand(compute_domain_createCmd)

	carapace.Gen(compute_domain_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Domain", "TTL").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
