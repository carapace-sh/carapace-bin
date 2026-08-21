package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_pool_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all backend storage pools known to the scheduler (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_pool_listCmd).Standalone()

	share_pool_listCmd.Flags().String("backend", "", "Filter results by backend name.")
	share_pool_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_pool_listCmd.Flags().Bool("detail", false, "Show detailed information about pools.")
	share_pool_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_pool_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_pool_listCmd.Flags().String("host", "", "Filter results by host name.")
	share_pool_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_pool_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_pool_listCmd.Flags().String("pool", "", "Filter results by pool name.")
	share_pool_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_pool_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_pool_listCmd.Flags().String("share-type", "", "Filter results by share type name or ID.")
	share_pool_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_pool_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_pool_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_poolCmd.AddCommand(share_pool_listCmd)
}
