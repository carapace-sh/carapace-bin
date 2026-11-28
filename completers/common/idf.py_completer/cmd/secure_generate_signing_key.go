package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureGenerateSigningKeyCmd = &cobra.Command{
	Use:   "secure-generate-signing-key",
	Short: "Generate secure boot signing key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureGenerateSigningKeyCmd).Standalone()
	secureGenerateSigningKeyCmd.Flags().StringP("scheme", "s", "", "Scheme of secure boot signing")
	secureGenerateSigningKeyCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	rootCmd.AddCommand(secureGenerateSigningKeyCmd)
	carapace.Gen(secureGenerateSigningKeyCmd).FlagCompletion(carapace.ActionMap{
		"scheme":  carapace.ActionValues("ecdsa256", "ecdsa384", "rsa3072"),
		"version": carapace.ActionValues("1", "2"),
	})
	carapace.Gen(secureGenerateSigningKeyCmd).PositionalCompletion(carapace.ActionFiles())
}
