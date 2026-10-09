package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List server migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_listCmd).Standalone()

	server_migration_listCmd.Flags().String("changes-before", "", "List only migrations changed earlier or equal to a certain point of time.")
	server_migration_listCmd.Flags().String("changes-since", "", "List only migrations changed later or equal to a certain point of time.")
	server_migration_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_migration_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_migration_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_migration_listCmd.Flags().String("host", "", "Filter migrations by source or destination host")
	server_migration_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	server_migration_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	server_migration_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	server_migration_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_migration_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_migration_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_migration_listCmd.Flags().String("project", "", "Filter migrations by project (name or ID) (supported with --os-compute-api-version 2.80 or above)")
	server_migration_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	server_migration_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	server_migration_listCmd.Flags().String("server", "", "Filter migrations by server (name or ID)")
	server_migration_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	server_migration_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	server_migration_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	server_migration_listCmd.Flags().String("status", "", "Filter migrations by status")
	server_migration_listCmd.Flags().String("type", "", "Filter migrations by type")
	server_migration_listCmd.Flags().String("user", "", "Filter migrations by user (name or ID) (supported with --os-compute-api-version 2.80 or above)")
	server_migration_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	server_migrationCmd.AddCommand(server_migration_listCmd)
}
