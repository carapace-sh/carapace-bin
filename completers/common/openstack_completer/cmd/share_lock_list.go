package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_lock_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all resource locks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_listCmd).Standalone()

	share_lock_listCmd.Flags().Bool("all-projects", false, "Filter resource locks for all projects.")
	share_lock_listCmd.Flags().String("before", "", "Filter resource locks created before given date.")
	share_lock_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_lock_listCmd.Flags().String("context", "", "Filter resource locks by context.")
	share_lock_listCmd.Flags().String("detailed", "", "Show detailed information about filtered resource locks.")
	share_lock_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_lock_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_lock_listCmd.Flags().String("id", "", "Filter resource locks by ID.")
	share_lock_listCmd.Flags().String("limit", "", "Number of resource locks to list.")
	share_lock_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_lock_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_lock_listCmd.Flags().String("offset", "", "Starting position of resource lock records in a paginated list.")
	share_lock_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_lock_listCmd.Flags().String("project", "", "Filter resource locks for specific project by name or ID, combine with --all-projects (Admin only).")
	share_lock_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_lock_listCmd.Flags().String("resource_action", "", "Filter resource locks by resource action.")
	share_lock_listCmd.Flags().String("resource_id", "", "Filter resource locks for a resource by ID, specify --resource-type to look up by name.")
	share_lock_listCmd.Flags().String("resource_type", "", "Filter resource locks by type of resource.")
	share_lock_listCmd.Flags().String("since", "", "Filter resource locks created since given date.")
	share_lock_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_lock_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_lock_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_lock_listCmd.Flags().String("sort_dir", "", "Sort direction, available values are ('asc', 'desc').")
	share_lock_listCmd.Flags().String("sort_key", "", "Key to be sorted, available keys are ('id', 'created_at', 'updated_at', 'resource_id', 'resource_type', 'resource_action', 'lock_reason').")
	share_lock_listCmd.Flags().String("user", "", "Filter resource locks for specific user by name or ID, combine with --all-projects to search across projects (Admin only).")
	share_lockCmd.AddCommand(share_lock_listCmd)
}
