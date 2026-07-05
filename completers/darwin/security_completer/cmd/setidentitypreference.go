package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setIdentityPreferenceCmd = &cobra.Command{
	Use:   "set-identity-preference",
	Short: "Set the preferred identity for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setIdentityPreferenceCmd).Standalone()
	setIdentityPreferenceCmd.Flags().StringP("hash", "Z", "", "Specify identity by SHA-1 hash of certificate")
	setIdentityPreferenceCmd.Flags().StringP("identity", "c", "", "Specify identity by common name of the certificate")
	setIdentityPreferenceCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (optional)")
	setIdentityPreferenceCmd.Flags().BoolP("no-identity", "n", false, "Specify no identity (clears existing preference)")
	setIdentityPreferenceCmd.Flags().StringP("service", "s", "", "Specify service (URL, RFC822 email, DNS host, or other)")
	rootCmd.AddCommand(setIdentityPreferenceCmd)
}
