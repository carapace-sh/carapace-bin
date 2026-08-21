package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_rbac_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List network RBAC policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rbac_listCmd).Standalone()

	network_rbac_listCmd.Flags().String("action", "", "List only network RBAC policies with the specified action (\"access_as_external\" or \"access_as_shared\")")
	network_rbac_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_rbac_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_rbac_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_rbac_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_rbac_listCmd.Flags().Bool("long", false, "List additional fields in output")
	network_rbac_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_rbac_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_rbac_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_rbac_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_rbac_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_rbac_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_rbac_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_rbac_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_rbac_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_rbac_listCmd.Flags().String("target-project", "", "List only network RBAC policies with the specified target project (name or ID)")
	network_rbac_listCmd.Flags().String("type", "", "List only network RBAC policies with the specified object type (\"address_group\", \"address_scope\", \"security_group\", \"subnetpool\", \"qos_policy\" or \"network\")")
	network_rbacCmd.AddCommand(network_rbac_listCmd)
}
