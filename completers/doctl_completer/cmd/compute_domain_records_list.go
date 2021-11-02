package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_records_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the DNS records for a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_records_listCmd).Standalone()
	compute_domain_records_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Type`, `Name`, `Data`, `Priority`, `Port`, `TTL`, `Weight`")
	compute_domain_records_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_domain_recordsCmd.AddCommand(compute_domain_records_listCmd)

	carapace.Gen(compute_domain_records_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Type", "Name", "Data", "Priority", "Port", "TTL", "Weight").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
