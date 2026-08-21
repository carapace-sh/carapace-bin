package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_service_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new tap service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_service_createCmd).Standalone()

	tap_service_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_service_createCmd.Flags().String("description", "", "Description of the tap service.")
	tap_service_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_service_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_service_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_service_createCmd.Flags().String("name", "", "Name of the tap service.")
	tap_service_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_service_createCmd.Flags().String("port", "", "Port (name or ID) to connect to the tap service.")
	tap_service_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_service_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_service_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	tap_service_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	tap_service_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_service_createCmd.MarkFlagRequired("port")
	tap_serviceCmd.AddCommand(tap_service_createCmd)
}
