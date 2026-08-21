package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_listCmd).Standalone()

	share_snapshot_listCmd.Flags().Bool("all-projects", false, "Display snapshots from all projects (Admin only).")
	share_snapshot_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_snapshot_listCmd.Flags().Bool("count", false, "The total count of share snapshots before pagination is applied.")
	share_snapshot_listCmd.Flags().String("description", "", "Filter results by description.")
	share_snapshot_listCmd.Flags().String("description~", "", "Filter results matching a share snapshot description pattern.")
	share_snapshot_listCmd.Flags().Bool("detail", false, "List share snapshots with details")
	share_snapshot_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_snapshot_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_snapshot_listCmd.Flags().String("limit", "", "Limit the number of snapshots returned")
	share_snapshot_listCmd.Flags().String("marker", "", "The last share ID of the previous page")
	share_snapshot_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_snapshot_listCmd.Flags().String("name", "", "Filter results by name.")
	share_snapshot_listCmd.Flags().String("name~", "", "Filter results matching a share snapshot name pattern.")
	share_snapshot_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_snapshot_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_snapshot_listCmd.Flags().String("property", "", "Filter snapshots having a given metadata key=value property.")
	share_snapshot_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_snapshot_listCmd.Flags().String("share", "", "Name or ID of a share to filter results by.")
	share_snapshot_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions(asc or desc) (default: name:asc), multiple keys and directions can be specified separated by comma")
	share_snapshot_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_snapshot_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_snapshot_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_snapshot_listCmd.Flags().String("status", "", "Filter results by status")
	share_snapshot_listCmd.Flags().String("usage", "", "Option to filter snapshots by usage.")
	share_snapshotCmd.AddCommand(share_snapshot_listCmd)
}
