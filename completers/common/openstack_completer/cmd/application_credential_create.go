package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var application_credential_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new application credential",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(application_credential_createCmd).Standalone()

	application_credential_createCmd.Flags().String("access-rules", "", "Either a string or file path containing a JSON-formatted list of access rules, each containing a request method, path, and service, for example '[{\"method\": \"GET\", \"path\": \"/v2.1/servers\", \"service\": \"compute\"}]'")
	application_credential_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	application_credential_createCmd.Flags().String("description", "", "Application credential description")
	application_credential_createCmd.Flags().String("expiration", "", "Sets an expiration date for the application credential, format of YYYY-mm-ddTHH:MM:SS (if not provided, the application credential will not expire)")
	application_credential_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	application_credential_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	application_credential_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	application_credential_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	application_credential_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	application_credential_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	application_credential_createCmd.Flags().Bool("restricted", false, "Prohibit application credential from creating and deleting other application credentials and trusts (this is the default behavior)")
	application_credential_createCmd.Flags().String("role", "", "Roles to authorize (name or ID) (repeat option to set multiple values)")
	application_credential_createCmd.Flags().String("secret", "", "Secret to use for authentication (if not provided, one will be generated)")
	application_credential_createCmd.Flags().Bool("unrestricted", false, "Enable application credential to create and delete other application credentials and trusts (this is potentially dangerous behavior and is disabled by default)")
	application_credential_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	application_credentialCmd.AddCommand(application_credential_createCmd)
}
