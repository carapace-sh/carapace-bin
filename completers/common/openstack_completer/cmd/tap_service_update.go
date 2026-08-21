package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_service_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a tap service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_service_updateCmd).Standalone()

	tap_service_updateCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_service_updateCmd.Flags().String("description", "", "Description of the tap service.")
	tap_service_updateCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_service_updateCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_service_updateCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_service_updateCmd.Flags().String("name", "", "Name of the tap service.")
	tap_service_updateCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_service_updateCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_service_updateCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_service_updateCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_serviceCmd.AddCommand(tap_service_updateCmd)
}
