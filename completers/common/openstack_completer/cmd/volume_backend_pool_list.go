package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backend_pool_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List pool command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backend_pool_listCmd).Standalone()

	volume_backend_pool_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_backend_pool_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_backend_pool_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_backend_pool_listCmd.Flags().Bool("long", false, "Show detailed information about pools.")
	volume_backend_pool_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_backend_pool_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_backend_pool_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_backend_pool_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_backend_pool_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_backend_pool_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_backend_pool_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_backend_poolCmd.AddCommand(volume_backend_pool_listCmd)
}
