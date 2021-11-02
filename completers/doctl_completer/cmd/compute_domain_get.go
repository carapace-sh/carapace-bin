package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_getCmd).Standalone()
	compute_domain_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Domain`, `TTL`")
	compute_domain_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_domainCmd.AddCommand(compute_domain_getCmd)

	carapace.Gen(compute_domain_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Domain", "TTL").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
