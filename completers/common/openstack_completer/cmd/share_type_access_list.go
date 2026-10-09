package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_access_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get access list for share type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_access_listCmd).Standalone()

	share_type_access_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_type_access_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_type_access_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_type_access_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_type_access_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_type_access_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_type_access_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_type_access_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_type_access_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_type_access_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_type_accessCmd.AddCommand(share_type_access_listCmd)
}
