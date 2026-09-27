package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display port details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_showCmd).Standalone()

	port_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	port_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	port_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	port_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	port_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	port_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	port_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	port_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	portCmd.AddCommand(port_showCmd)
}
