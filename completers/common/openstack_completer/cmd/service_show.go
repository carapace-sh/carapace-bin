package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display service details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_showCmd).Standalone()

	service_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	service_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	serviceCmd.AddCommand(service_showCmd)
}
