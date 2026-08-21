package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share group snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_listCmd).Standalone()

	share_group_snapshot_listCmd.Flags().Bool("all-projects", false, "Display information from all projects (Admin only).")
	share_group_snapshot_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_snapshot_listCmd.Flags().Bool("detailed", false, "Show detailed information about share group snapshot. ")
	share_group_snapshot_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_snapshot_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_snapshot_listCmd.Flags().String("limit", "", "Limit the number of share groups returned")
	share_group_snapshot_listCmd.Flags().String("marker", "", "The last share group snapshot ID of the previous page")
	share_group_snapshot_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_snapshot_listCmd.Flags().String("name", "", "Filter results by name.")
	share_group_snapshot_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_snapshot_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_snapshot_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_group_snapshot_listCmd.Flags().String("share-group", "", "Filter results by share group name or ID.")
	share_group_snapshot_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions(asc or desc) (default: name:asc), multiple keys and directions can be specified separated by comma")
	share_group_snapshot_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_group_snapshot_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_group_snapshot_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_group_snapshot_listCmd.Flags().String("status", "", "Filter results by status.")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_listCmd)
}
