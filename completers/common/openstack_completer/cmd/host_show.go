package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var host_showCmd = &cobra.Command{
	Use:   "show",
	Short: "DEPRECATED: Display host details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(host_showCmd).Standalone()

	host_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	host_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	host_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	host_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	host_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	host_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	host_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	host_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	host_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	host_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	hostCmd.AddCommand(host_showCmd)
}
