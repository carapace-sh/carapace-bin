package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ip_availability_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List IP availability for network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ip_availability_listCmd).Standalone()

	ip_availability_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	ip_availability_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	ip_availability_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	ip_availability_listCmd.Flags().String("ip-version", "", "List only IP availability with the specified IP version networks (4 or 6, default is 4)")
	ip_availability_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	ip_availability_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	ip_availability_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	ip_availability_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	ip_availability_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	ip_availability_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	ip_availability_listCmd.Flags().String("project", "", "List only IP availability with the specified project (name or ID)")
	ip_availability_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	ip_availability_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	ip_availability_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	ip_availability_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	ip_availability_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	ip_availabilityCmd.AddCommand(ip_availability_listCmd)
}
