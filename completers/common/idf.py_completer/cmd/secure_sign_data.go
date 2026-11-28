package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureSignDataCmd = &cobra.Command{
	Use:   "secure-sign-data",
	Short: "Sign data for secure boot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureSignDataCmd).Standalone()
	secureSignDataCmd.Flags().BoolP("append-signatures", "a", false, "Append signature block(s) to already signed image")
	secureSignDataCmd.Flags().StringP("keyfile", "k", "", "Private key file for signing")
	secureSignDataCmd.Flags().StringP("output", "o", "", "Output file for signed digest image")
	secureSignDataCmd.Flags().String("pub-key", "", "Public key files corresponding to the private key")
	secureSignDataCmd.Flags().String("signature", "", "Pre-calculated signatures")
	secureSignDataCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	carapace.Gen(secureSignDataCmd).FlagCompletion(carapace.ActionMap{
		"keyfile":   carapace.ActionFiles(),
		"output":    carapace.ActionFiles(),
		"pub-key":   carapace.ActionFiles(),
		"signature": carapace.ActionFiles(),
		"version":   carapace.ActionValues("1", "2"),
	})
	carapace.Gen(secureSignDataCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureSignDataCmd)
}
