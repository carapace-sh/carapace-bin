package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getIdentityPreferenceCmd = &cobra.Command{
	Use:   "get-identity-preference",
	Short: "Get the preferred identity for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getIdentityPreferenceCmd).Standalone()
	getIdentityPreferenceCmd.Flags().BoolP("common-name", "c", false, "Print common name of the preferred identity certificate")
	getIdentityPreferenceCmd.Flags().BoolP("hash", "Z", false, "Print SHA-1 hash of the preferred identity certificate")
	getIdentityPreferenceCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (optional)")
	getIdentityPreferenceCmd.Flags().BoolP("pem", "p", false, "Output identity certificate in PEM format")
	getIdentityPreferenceCmd.Flags().StringP("service", "s", "", "Specify service (URL, RFC822 email, DNS host, or other)")
	rootCmd.AddCommand(getIdentityPreferenceCmd)
}
