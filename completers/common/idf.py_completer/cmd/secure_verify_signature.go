package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureVerifySignatureCmd = &cobra.Command{
	Use:   "secure-verify-signature",
	Short: "Verify secure boot signature",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureVerifySignatureCmd).Standalone()

	secureVerifySignatureCmd.Flags().StringP("keyfile", "k", "", "Public key file for verification")
	secureVerifySignatureCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	rootCmd.AddCommand(secureVerifySignatureCmd)

	carapace.Gen(secureVerifySignatureCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"version": carapace.ActionValues("1", "2"),
	})

	carapace.Gen(secureVerifySignatureCmd).PositionalCompletion(carapace.ActionFiles())
}
