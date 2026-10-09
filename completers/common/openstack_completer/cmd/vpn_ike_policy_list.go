package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List IKE policies that belong to a given project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policy_listCmd).Standalone()

	vpn_ike_policy_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ike_policy_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ike_policy_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ike_policy_listCmd.Flags().Bool("long", false, "List additional fields in output")
	vpn_ike_policy_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ike_policy_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ike_policy_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ike_policy_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	vpn_ike_policy_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	vpn_ike_policy_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	vpn_ike_policy_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	vpn_ike_policyCmd.AddCommand(vpn_ike_policy_listCmd)
}
