package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network trunks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_listCmd).Standalone()

	network_trunk_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_trunk_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_trunk_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_trunk_listCmd.Flags().Bool("long", false, "List additional fields in output")
	network_trunk_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_trunk_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_trunk_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_trunk_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_trunk_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_trunk_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_trunk_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_trunkCmd.AddCommand(network_trunk_listCmd)
}
