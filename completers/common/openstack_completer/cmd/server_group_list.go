package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all server groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_group_listCmd).Standalone()

	server_group_listCmd.Flags().Bool("all-projects", false, "Display information from all projects (admin only)")
	server_group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_group_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	server_group_listCmd.Flags().Bool("long", false, "List additional fields in output")
	server_group_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	server_group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_group_listCmd.Flags().String("offset", "", "The (zero-based) offset of the first item in the collection to return.")
	server_group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	server_group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	server_group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	server_group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	server_groupCmd.AddCommand(server_group_listCmd)
}
