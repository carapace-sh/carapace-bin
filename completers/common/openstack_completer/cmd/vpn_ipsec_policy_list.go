package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List IPsec policies that belong to a given project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policy_listCmd).Standalone()

	vpn_ipsec_policy_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ipsec_policy_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ipsec_policy_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ipsec_policy_listCmd.Flags().Bool("long", false, "List additional fields in output")
	vpn_ipsec_policy_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ipsec_policy_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ipsec_policy_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ipsec_policy_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	vpn_ipsec_policy_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	vpn_ipsec_policy_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	vpn_ipsec_policy_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	vpn_ipsec_policyCmd.AddCommand(vpn_ipsec_policy_listCmd)
}
