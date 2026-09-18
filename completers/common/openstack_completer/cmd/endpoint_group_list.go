package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List endpoint groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_group_listCmd).Standalone()

	endpoint_group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	endpoint_group_listCmd.Flags().String("domain", "", "Domain owning <project> (name or ID)")
	endpoint_group_listCmd.Flags().String("endpointgroup", "", "Endpoint Group (name or ID)")
	endpoint_group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	endpoint_group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	endpoint_group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	endpoint_group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	endpoint_group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	endpoint_group_listCmd.Flags().String("project", "", "Project (name or ID)")
	endpoint_group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	endpoint_group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	endpoint_group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	endpoint_group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	endpoint_groupCmd.AddCommand(endpoint_group_listCmd)
}
