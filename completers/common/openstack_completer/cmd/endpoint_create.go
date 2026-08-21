package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new endpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_createCmd).Standalone()

	endpoint_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	endpoint_createCmd.Flags().Bool("disable", false, "Disable endpoint")
	endpoint_createCmd.Flags().Bool("enable", false, "Enable endpoint (default)")
	endpoint_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	endpoint_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	endpoint_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	endpoint_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	endpoint_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	endpoint_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	endpoint_createCmd.Flags().String("region", "", "New endpoint region ID")
	endpoint_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	endpointCmd.AddCommand(endpoint_createCmd)
}
