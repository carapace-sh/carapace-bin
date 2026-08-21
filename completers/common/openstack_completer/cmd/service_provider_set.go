package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_provider_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set service provider properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_provider_setCmd).Standalone()

	service_provider_setCmd.Flags().String("auth-url", "", "New Authentication URL of remote federated service provider")
	service_provider_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_provider_setCmd.Flags().String("description", "", "New service provider description")
	service_provider_setCmd.Flags().Bool("disable", false, "Disable the service provider")
	service_provider_setCmd.Flags().Bool("enable", false, "Enable the service provider")
	service_provider_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_provider_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_provider_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_provider_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_provider_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	service_provider_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_provider_setCmd.Flags().String("service-provider-url", "", "New service provider URL, where SAML assertions are sent")
	service_provider_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	service_providerCmd.AddCommand(service_provider_setCmd)
}
