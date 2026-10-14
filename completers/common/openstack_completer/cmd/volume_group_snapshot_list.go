package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all volume group snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_snapshot_listCmd).Standalone()

	volume_group_snapshot_listCmd.Flags().Bool("all-projects", false, "Shows details for all projects (admin only).")
	volume_group_snapshot_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_snapshot_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_snapshot_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_snapshot_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_snapshot_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_snapshot_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_snapshot_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_group_snapshot_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_group_snapshot_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_group_snapshot_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_group_snapshotCmd.AddCommand(volume_group_snapshot_listCmd)
}
