package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_flow_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a tap flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_flow_updateCmd).Standalone()

	tap_flow_updateCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_flow_updateCmd.Flags().String("description", "", "Description of the tap flow.")
	tap_flow_updateCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_flow_updateCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_flow_updateCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_flow_updateCmd.Flags().String("name", "", "Name of the tap flow.")
	tap_flow_updateCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_flow_updateCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_flow_updateCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_flow_updateCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_flowCmd.AddCommand(tap_flow_updateCmd)
}
