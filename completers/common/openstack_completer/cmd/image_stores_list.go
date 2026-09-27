package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_stores_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get available backends (only valid with Multi-Backend support)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_stores_listCmd).Standalone()

	image_stores_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_stores_listCmd.Flags().Bool("detail", false, "Shows details of stores (admin only) (requires --os-image-api-version 2.15 or later)")
	image_stores_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_stores_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_stores_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_stores_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_stores_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_stores_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	image_stores_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	image_stores_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	image_stores_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	image_storesCmd.AddCommand(image_stores_listCmd)
}
