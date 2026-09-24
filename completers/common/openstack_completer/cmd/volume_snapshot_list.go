package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List volume snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_snapshot_listCmd).Standalone()

	volume_snapshot_listCmd.Flags().Bool("all-projects", false, "Include all projects (admin only)")
	volume_snapshot_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_snapshot_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_snapshot_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_snapshot_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	volume_snapshot_listCmd.Flags().Bool("long", false, "List additional fields in output")
	volume_snapshot_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	volume_snapshot_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	volume_snapshot_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_snapshot_listCmd.Flags().String("name", "", "Filters results by a name.")
	volume_snapshot_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_snapshot_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_snapshot_listCmd.Flags().String("project", "", "Filter results by project (name or ID) (admin only)")
	volume_snapshot_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	volume_snapshot_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_snapshot_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_snapshot_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_snapshot_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_snapshot_listCmd.Flags().String("status", "", "Filters results by a status.")
	volume_snapshot_listCmd.Flags().String("volume", "", "Filters results by a volume (name or ID).")
	volume_snapshotCmd.AddCommand(volume_snapshot_listCmd)
}
