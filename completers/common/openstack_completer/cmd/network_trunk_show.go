package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information of a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_showCmd).Standalone()

	network_trunk_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_trunk_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_trunk_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_trunk_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_trunk_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_trunk_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_trunk_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_trunk_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_trunkCmd.AddCommand(network_trunk_showCmd)
}
