package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var host_listCmd = &cobra.Command{
	Use:   "list",
	Short: "DEPRECATED: List hosts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(host_listCmd).Standalone()

	host_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	host_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	host_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	host_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	host_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	host_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	host_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	host_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	host_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	host_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	host_listCmd.Flags().String("zone", "", "Only return hosts in the availability zone")
	hostCmd.AddCommand(host_listCmd)
}
