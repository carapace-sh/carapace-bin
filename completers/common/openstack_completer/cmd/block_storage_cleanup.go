package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Do block storage cleanup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_cleanupCmd).Standalone()

	block_storage_cleanupCmd.Flags().String("binary", "", "Name of the service binary.")
	block_storage_cleanupCmd.Flags().String("cluster", "", "Name of block storage cluster in which cleanup needs to be performed (name only)")
	block_storage_cleanupCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	block_storage_cleanupCmd.Flags().Bool("disabled", false, "Filter by disabled status.")
	block_storage_cleanupCmd.Flags().Bool("down", false, "Filter by down status.")
	block_storage_cleanupCmd.Flags().Bool("enabled", false, "Filter by enabled status.")
	block_storage_cleanupCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	block_storage_cleanupCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	block_storage_cleanupCmd.Flags().String("host", "", "Host where the service resides.")
	block_storage_cleanupCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	block_storage_cleanupCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	block_storage_cleanupCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	block_storage_cleanupCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	block_storage_cleanupCmd.Flags().String("resource-id", "", "UUID of a resource to cleanup.")
	block_storage_cleanupCmd.Flags().String("resource-type", "", "Type of resource to cleanup.")
	block_storage_cleanupCmd.Flags().String("service-id", "", "The service ID field from the DB, not the UUID of the service.")
	block_storage_cleanupCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	block_storage_cleanupCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	block_storage_cleanupCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	block_storage_cleanupCmd.Flags().Bool("up", false, "Filter by up status.")
	block_storageCmd.AddCommand(block_storage_cleanupCmd)
}
