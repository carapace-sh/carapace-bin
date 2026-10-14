package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_flow_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tap flow details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_flow_showCmd).Standalone()

	tap_flow_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_flow_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_flow_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_flow_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_flow_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_flow_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_flow_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_flow_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_flowCmd.AddCommand(tap_flow_showCmd)
}
