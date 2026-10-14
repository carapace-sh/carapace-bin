package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_createCmd).Standalone()

	service_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_createCmd.Flags().String("description", "", "New service description")
	service_createCmd.Flags().Bool("disable", false, "Disable service")
	service_createCmd.Flags().Bool("enable", false, "Enable service (default)")
	service_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_createCmd.Flags().String("name", "", "New service name")
	service_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	service_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	serviceCmd.AddCommand(service_createCmd)
}
