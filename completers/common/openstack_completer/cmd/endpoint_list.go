package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List endpoints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_listCmd).Standalone()

	endpoint_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	endpoint_listCmd.Flags().String("endpoint", "", "Endpoint to list filters")
	endpoint_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	endpoint_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	endpoint_listCmd.Flags().String("interface", "", "Filter by interface type (admin, public or internal)")
	endpoint_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	endpoint_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	endpoint_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	endpoint_listCmd.Flags().String("project", "", "Project to list filters (name or ID)")
	endpoint_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	endpoint_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	endpoint_listCmd.Flags().String("region", "", "Filter by region ID")
	endpoint_listCmd.Flags().String("service", "", "Filter by service (type, name or ID)")
	endpoint_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	endpoint_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	endpoint_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	endpointCmd.AddCommand(endpoint_listCmd)
}
