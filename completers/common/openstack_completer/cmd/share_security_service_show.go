package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show security service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_showCmd).Standalone()

	share_security_service_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_security_service_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_security_service_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_security_service_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_security_service_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_security_service_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_security_service_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_security_service_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_security_serviceCmd.AddCommand(share_security_service_showCmd)
}
