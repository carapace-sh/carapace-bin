package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var rsaCmd = &cobra.Command{
	Use:     "rsa",
	Short:   "RSA key management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rsaCmd).Standalone()

	rsaCmd.Flags().BoolS("RSAPublicKey_in", "RSAPublicKey_in", false, "Input is an RSAPublicKey")
	rsaCmd.Flags().BoolS("RSAPublicKey_out", "RSAPublicKey_out", false, "Output is an RSAPublicKey")
	rsaCmd.Flags().BoolS("check", "check", false, "Verify key consistency")
	rsaCmd.Flags().StringS("in", "in", "", "Input file")
	rsaCmd.Flags().StringS("inform", "inform", "", "Input format (DER/PEM/P12)")
	rsaCmd.Flags().BoolS("modulus", "modulus", false, "Print the RSA key modulus")
	rsaCmd.Flags().BoolS("noout", "noout", false, "Don't print key out")
	rsaCmd.Flags().StringS("out", "out", "", "Output file")
	rsaCmd.Flags().StringS("outform", "outform", "", "Output format, one of DER PEM PVK")
	rsaCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	rsaCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	rsaCmd.Flags().BoolS("pubin", "pubin", false, "Expect a public key in input file")
	rsaCmd.Flags().BoolS("pubout", "pubout", false, "Output a public key")
	rsaCmd.Flags().BoolS("pvk-none", "pvk-none", false, "Don't enforce PVK encoding")
	rsaCmd.Flags().BoolS("pvk-strong", "pvk-strong", false, "Enable 'Strong' PVK encoding level (default)")
	rsaCmd.Flags().BoolS("pvk-weak", "pvk-weak", false, "Enable 'Weak' PVK encoding level")
	rsaCmd.Flags().BoolS("text", "text", false, "Print the key in text")
	rsaCmd.Flags().BoolS("traditional", "traditional", false, "Use traditional format for private keys")
	common.AddProviderFlags(rsaCmd)
	rootCmd.AddCommand(rsaCmd)

	carapace.Gen(rsaCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM", "P12"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
