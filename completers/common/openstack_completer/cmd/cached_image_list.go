package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cached_image_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get Cache State",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cached_image_listCmd).Standalone()

	cached_image_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	cached_image_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	cached_image_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	cached_image_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	cached_image_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	cached_image_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	cached_image_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	cached_image_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	cached_image_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	cached_image_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	cached_imageCmd.AddCommand(cached_image_listCmd)
}
