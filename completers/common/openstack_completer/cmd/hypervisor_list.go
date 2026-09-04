package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hypervisor_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List hypervisors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hypervisor_listCmd).Standalone()

	hypervisor_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	hypervisor_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	hypervisor_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	hypervisor_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	hypervisor_listCmd.Flags().Bool("long", false, "List additional fields in output")
	hypervisor_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	hypervisor_listCmd.Flags().String("matching", "", "Filter hypervisors using <hostname> substringHypervisor Type and Host IP are not returned when using microversion 2.52 or lower")
	hypervisor_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	hypervisor_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	hypervisor_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	hypervisor_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	hypervisor_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	hypervisor_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	hypervisor_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	hypervisor_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	hypervisorCmd.AddCommand(hypervisor_listCmd)
}
