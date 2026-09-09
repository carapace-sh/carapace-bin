package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_provider_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new service provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_provider_createCmd).Standalone()

	service_provider_createCmd.Flags().String("auth-url", "", "Authentication URL of remote federated service provider (required)")
	service_provider_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_provider_createCmd.Flags().String("description", "", "New service provider description")
	service_provider_createCmd.Flags().Bool("disable", false, "Disable the service provider")
	service_provider_createCmd.Flags().Bool("enable", false, "Enable the service provider (default)")
	service_provider_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_provider_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_provider_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_provider_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_provider_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	service_provider_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_provider_createCmd.Flags().String("service-provider-url", "", "A service URL where SAML assertions are being sent (required)")
	service_provider_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	service_provider_createCmd.MarkFlagRequired("auth-url")
	service_provider_createCmd.MarkFlagRequired("service-provider-url")
	service_providerCmd.AddCommand(service_provider_createCmd)
}
