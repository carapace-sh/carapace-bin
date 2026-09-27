package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display network flavor details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_showCmd).Standalone()

	network_flavor_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_flavor_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_flavor_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_flavor_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_flavor_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_flavor_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_flavor_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_flavor_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_flavorCmd.AddCommand(network_flavor_showCmd)
}
