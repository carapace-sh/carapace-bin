package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_provider_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new identity provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_createCmd).Standalone()

	identity_provider_createCmd.Flags().String("authorization-ttl", "", "Time to keep the role assignments for users authenticating via this identity provider.")
	identity_provider_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	identity_provider_createCmd.Flags().String("description", "", "New identity provider description")
	identity_provider_createCmd.Flags().Bool("disable", false, "Disable the identity provider")
	identity_provider_createCmd.Flags().String("domain", "", "Domain to associate with the identity provider.")
	identity_provider_createCmd.Flags().Bool("enable", false, "Enable identity provider (default)")
	identity_provider_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	identity_provider_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	identity_provider_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	identity_provider_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	identity_provider_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	identity_provider_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	identity_provider_createCmd.Flags().String("remote-id", "", "Remote IDs to associate with the Identity Provider (repeat option to provide multiple values)")
	identity_provider_createCmd.Flags().String("remote-id-file", "", "Name of a file that contains many remote IDs to associate with the identity provider, one per line")
	identity_provider_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	identity_providerCmd.AddCommand(identity_provider_createCmd)
}
