package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dsaCmd = &cobra.Command{
	Use:     "dsa",
	Short:   "DSA Data Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsaCmd).Standalone()

	dsaCmd.Flags().StringS("in", "in", "", "Input key")
	dsaCmd.Flags().StringS("inform", "inform", "", "Input format (DER/PEM/PVK); has no effect")
	dsaCmd.Flags().BoolS("modulus", "modulus", false, "Print the DSA public value")
	dsaCmd.Flags().BoolS("noout", "noout", false, "Don't print key out")
	dsaCmd.Flags().StringS("out", "out", "", "Output file")
	dsaCmd.Flags().StringS("outform", "outform", "", "Output format, DER PEM PVK")
	dsaCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	dsaCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	dsaCmd.Flags().BoolS("pubin", "pubin", false, "Expect a public key in input file")
	dsaCmd.Flags().BoolS("pubout", "pubout", false, "Output public key, not private")
	dsaCmd.Flags().BoolS("pvk-none", "pvk-none", false, "Don't enforce PVK encoding")
	dsaCmd.Flags().BoolS("pvk-strong", "pvk-strong", false, "Enable 'Strong' PVK encoding level (default)")
	dsaCmd.Flags().BoolS("pvk-weak", "pvk-weak", false, "Enable 'Weak' PVK encoding level")
	dsaCmd.Flags().BoolS("text", "text", false, "Print the key in text")
	common.AddProviderFlags(dsaCmd)
	rootCmd.AddCommand(dsaCmd)

	carapace.Gen(dsaCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
