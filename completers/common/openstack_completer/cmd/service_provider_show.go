package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_provider_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display service provider details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_provider_showCmd).Standalone()

	service_provider_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_provider_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_provider_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_provider_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_provider_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_provider_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	service_provider_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_provider_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	service_providerCmd.AddCommand(service_provider_showCmd)
}
