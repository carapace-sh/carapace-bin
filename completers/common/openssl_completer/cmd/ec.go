package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var ecCmd = &cobra.Command{
	Use:     "ec",
	Short:   "EC (Elliptic curve) key processing",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecCmd).Standalone()

	ecCmd.Flags().BoolS("check", "check", false, "check key consistency")
	ecCmd.Flags().StringS("conv_form", "conv_form", "", "Specifies the point conversion form")
	ecCmd.Flags().StringS("in", "in", "", "Input file")
	ecCmd.Flags().StringS("inform", "inform", "", "Input format (DER/PEM/P12)")
	ecCmd.Flags().BoolS("no_public", "no_public", false, "exclude public key from private key")
	ecCmd.Flags().BoolS("noout", "noout", false, "Don't print key out")
	ecCmd.Flags().StringS("out", "out", "", "Output file")
	ecCmd.Flags().StringS("outform", "outform", "", "Output format - DER or PEM")
	ecCmd.Flags().StringS("param_enc", "param_enc", "", "Specifies the way the ec parameters are encoded")
	ecCmd.Flags().BoolS("param_out", "param_out", false, "Print the elliptic curve parameters")
	ecCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	ecCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	ecCmd.Flags().BoolS("pubin", "pubin", false, "Expect a public key in input file")
	ecCmd.Flags().BoolS("pubout", "pubout", false, "Output public key, not private")
	ecCmd.Flags().BoolS("text", "text", false, "Print the key")
	common.AddProviderFlags(ecCmd)
	rootCmd.AddCommand(ecCmd)

	carapace.Gen(ecCmd).FlagCompletion(carapace.ActionMap{
		"conv_form": carapace.ActionValues("compressed", "hybrid", "uncompressed"),
		"in":        carapace.ActionFiles(),
		"inform":    carapace.ActionValues("DER", "PEM", "P12"),
		"out":       carapace.ActionFiles(),
		"outform":   carapace.ActionValues("DER", "PEM"),
		"param_enc": carapace.ActionValues("explicit", "named_curve"),
	})
}
