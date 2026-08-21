package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_policy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List QoS policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_policy_listCmd).Standalone()

	network_qos_policy_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_qos_policy_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_qos_policy_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_qos_policy_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_qos_policy_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_qos_policy_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_qos_policy_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_qos_policy_listCmd.Flags().Bool("no-share", false, "List only QoS policies not shared between projects")
	network_qos_policy_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_qos_policy_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_qos_policy_listCmd.Flags().String("project", "", "List only QoS policies with the specified project (name or ID)")
	network_qos_policy_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_qos_policy_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_qos_policy_listCmd.Flags().Bool("share", false, "List only QoS policies shared between projects")
	network_qos_policy_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_qos_policy_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_qos_policy_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_qos_policyCmd.AddCommand(network_qos_policy_listCmd)
}
