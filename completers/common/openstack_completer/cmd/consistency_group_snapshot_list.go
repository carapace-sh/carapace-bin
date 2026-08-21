package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List consistency group snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_snapshot_listCmd).Standalone()

	consistency_group_snapshot_listCmd.Flags().Bool("all-projects", false, "Show detail for all projects (admin only) (defaults to False)")
	consistency_group_snapshot_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consistency_group_snapshot_listCmd.Flags().String("consistency-group", "", "Filters results by a consistency group (name or ID)")
	consistency_group_snapshot_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consistency_group_snapshot_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consistency_group_snapshot_listCmd.Flags().Bool("long", false, "List additional fields in output")
	consistency_group_snapshot_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consistency_group_snapshot_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consistency_group_snapshot_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consistency_group_snapshot_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	consistency_group_snapshot_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	consistency_group_snapshot_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	consistency_group_snapshot_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	consistency_group_snapshot_listCmd.Flags().String("status", "", "Filters results by a status (\"available\", \"error\", \"creating\", \"deleting\" or \"error_deleting\")")
	consistency_group_snapshotCmd.AddCommand(consistency_group_snapshot_listCmd)
}
