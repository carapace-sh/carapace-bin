package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var rsautlCmd = &cobra.Command{
	Use:     "rsautl",
	Short:   "RSA command for signing, verification, encryption, and decryption",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rsautlCmd).Standalone()

	rsautlCmd.Flags().BoolS("asn1parse", "asn1parse", false, "Run output through asn1parse; useful with -verify")
	rsautlCmd.Flags().BoolS("certin", "certin", false, "Input is a cert carrying an RSA public key")
	rsautlCmd.Flags().BoolS("decrypt", "decrypt", false, "Decrypt with private key")
	rsautlCmd.Flags().BoolS("encrypt", "encrypt", false, "Encrypt with public key")
	rsautlCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	rsautlCmd.Flags().BoolS("hexdump", "hexdump", false, "Hex dump output")
	rsautlCmd.Flags().StringS("in", "in", "", "Input file")
	rsautlCmd.Flags().StringS("inkey", "inkey", "", "Input key, by default an RSA private key")
	rsautlCmd.Flags().StringS("keyform", "keyform", "", "Private key format (ENGINE, other values ignored)")
	rsautlCmd.Flags().BoolS("oaep", "oaep", false, "Use PKCS#1 OAEP")
	rsautlCmd.Flags().StringS("out", "out", "", "Output file")
	rsautlCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	rsautlCmd.Flags().BoolS("pkcs", "pkcs", false, "Use PKCS#1 v1.5 padding (default)")
	rsautlCmd.Flags().BoolS("pubin", "pubin", false, "Input key is an RSA public pkey")
	rsautlCmd.Flags().BoolS("raw", "raw", false, "Use no padding")
	rsautlCmd.Flags().BoolS("rev", "rev", false, "Reverse the order of the input buffer")
	rsautlCmd.Flags().BoolS("sign", "sign", false, "Sign with private key")
	rsautlCmd.Flags().BoolS("verify", "verify", false, "Verify with public key")
	rsautlCmd.Flags().BoolS("x931", "x931", false, "Use ANSI X9.31 padding")
	common.AddProviderFlags(rsautlCmd)
	common.AddRandomStateFlags(rsautlCmd)
	rootCmd.AddCommand(rsautlCmd)

	carapace.Gen(rsautlCmd).FlagCompletion(carapace.ActionMap{
		"engine":  action.ActionEngines(),
		"in":      carapace.ActionFiles(),
		"inkey":   carapace.ActionFiles(),
		"keyform": carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"out":     carapace.ActionFiles(),
	})
}
