package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_subport_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all subports for a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_subport_listCmd).Standalone()

	network_subport_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_subport_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_subport_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_subport_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_subport_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_subport_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_subport_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_subport_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_subport_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_subport_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_subport_listCmd.Flags().String("trunk", "", "List only subports belonging to this trunk (name or ID)")
	network_subport_listCmd.MarkFlagRequired("trunk")
	network_subportCmd.AddCommand(network_subport_listCmd)
}
