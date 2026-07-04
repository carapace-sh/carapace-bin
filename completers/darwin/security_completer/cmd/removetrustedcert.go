package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeTrustedCertCmd = &cobra.Command{
	Use:   "remove-trusted-cert",
	Short: "Remove trusted certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeTrustedCertCmd).Standalone()
	removeTrustedCertCmd.Flags().BoolP("admin", "d", false, "Remove from admin cert store (default: user)")
	removeTrustedCertCmd.Flags().BoolP("default", "D", false, "Remove Default Root Cert setting instead of actual cert setting")
	rootCmd.AddCommand(removeTrustedCertCmd)
	carapace.Gen(removeTrustedCertCmd).PositionalCompletion(carapace.ActionFiles())
}
