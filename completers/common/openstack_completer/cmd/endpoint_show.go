package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display endpoint details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_showCmd).Standalone()

	endpoint_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	endpoint_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	endpoint_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	endpoint_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	endpoint_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	endpoint_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	endpoint_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	endpoint_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	endpointCmd.AddCommand(endpoint_showCmd)
}
