package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_snapshot_manageable_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List manageable snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_snapshot_manageable_listCmd).Standalone()

	block_storage_snapshot_manageable_listCmd.Flags().String("cluster", "", "Cinder cluster on which to list manageable snapshots.")
	block_storage_snapshot_manageable_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	block_storage_snapshot_manageable_listCmd.Flags().String("detailed", "", "==SUPPRESS==")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	block_storage_snapshot_manageable_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	block_storage_snapshot_manageable_listCmd.Flags().String("limit", "", "Maximum number of snapshots to return.")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("long", false, "List additional fields in output")
	block_storage_snapshot_manageable_listCmd.Flags().String("marker", "", "Begin returning snapshots that appear later in the snapshot list than that represented by this reference.")
	block_storage_snapshot_manageable_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	block_storage_snapshot_manageable_listCmd.Flags().String("offset", "", "Number of snapshots to skip after marker.")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	block_storage_snapshot_manageable_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	block_storage_snapshot_manageable_listCmd.Flags().String("sort", "", "Comma-separated list of sort keys and directions in the form of <key>[:<asc|desc>].")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	block_storage_snapshot_manageable_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	block_storage_snapshot_manageable_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	block_storage_snapshot_manageableCmd.AddCommand(block_storage_snapshot_manageable_listCmd)
}
