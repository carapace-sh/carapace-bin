package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkeyCmd = &cobra.Command{
	Use:     "pkey",
	Short:   "Public and private key management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkeyCmd).Standalone()

	pkeyCmd.Flags().BoolS("check", "check", false, "Check key consistency")
	pkeyCmd.Flags().StringS("ec_conv_form", "ec_conv_form", "", "Specifies the EC point conversion form in the encoding")
	pkeyCmd.Flags().StringS("ec_param_enc", "ec_param_enc", "", "Specifies the way the EC parameters are encoded")
	pkeyCmd.Flags().StringSliceS("encopt", "encopt", nil, "Private key encoder parameter")
	pkeyCmd.Flags().StringS("in", "in", "", "Input key")
	pkeyCmd.Flags().StringS("inform", "inform", "", "Key input format (DER/PEM)")
	pkeyCmd.Flags().BoolS("noout", "noout", false, "Do not output the key in encoded form")
	pkeyCmd.Flags().StringS("out", "out", "", "Output file for encoded and/or text output")
	pkeyCmd.Flags().StringS("outform", "outform", "", "Output encoding format (DER or PEM)")
	pkeyCmd.Flags().StringS("passin", "passin", "", "Key input pass phrase source")
	pkeyCmd.Flags().StringS("passout", "passout", "", "Output PEM file pass phrase source")
	pkeyCmd.Flags().BoolS("pubcheck", "pubcheck", false, "Check public key consistency")
	pkeyCmd.Flags().BoolS("pubin", "pubin", false, "Read only public components from key input")
	pkeyCmd.Flags().BoolS("pubout", "pubout", false, "Restrict encoded output to public components")
	pkeyCmd.Flags().BoolS("text", "text", false, "Output key components in plaintext")
	pkeyCmd.Flags().BoolS("text_pub", "text_pub", false, "Output only public key components in text form")
	pkeyCmd.Flags().BoolS("traditional", "traditional", false, "Use traditional format for private key PEM output")
	common.AddProviderFlags(pkeyCmd)
	rootCmd.AddCommand(pkeyCmd)

	carapace.Gen(pkeyCmd).FlagCompletion(carapace.ActionMap{
		"ec_conv_form": carapace.ActionValues("compressed", "hybrid", "uncompressed"),
		"ec_param_enc": carapace.ActionValues("named_curve", "explicit"),
		"in":           carapace.ActionFiles(),
		"inform":       carapace.ActionValues("DER", "PEM", "P12"),
		"out":          carapace.ActionFiles(),
		"outform":      carapace.ActionValues("DER", "PEM"),
	})
}
