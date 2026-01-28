package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var ecparamCmd = &cobra.Command{
	Use:     "ecparam",
	Short:   "EC parameter manipulation and generation",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecparamCmd).Standalone()

	ecparamCmd.Flags().BoolS("check", "check", false, "Validate the ec parameters")
	ecparamCmd.Flags().BoolS("check_named", "check_named", false, "Check that named EC curve parameters have not been modified")
	ecparamCmd.Flags().StringS("conv_form", "conv_form", "", "Specifies the point conversion form")
	ecparamCmd.Flags().BoolS("genkey", "genkey", false, "Generate ec key")
	ecparamCmd.Flags().StringS("in", "in", "", "Input file  - default stdin")
	ecparamCmd.Flags().StringS("inform", "inform", "", "Input format - default PEM (DER or PEM)")
	ecparamCmd.Flags().BoolS("list_curves", "list_curves", false, "Prints a list of all curve 'short names'")
	ecparamCmd.Flags().StringS("name", "name", "", "Use the ec parameters with specified 'short name'")
	ecparamCmd.Flags().BoolS("no_seed", "no_seed", false, "If 'explicit' parameters are chosen do not use the seed")
	ecparamCmd.Flags().BoolS("noout", "noout", false, "Do not print the ec parameter")
	ecparamCmd.Flags().StringS("out", "out", "", "Output file - default stdout")
	ecparamCmd.Flags().StringS("outform", "outform", "", "Output format - default PEM")
	ecparamCmd.Flags().StringS("param_enc", "param_enc", "", "Specifies the way the ec parameters are encoded")
	ecparamCmd.Flags().BoolS("text", "text", false, "Print the ec parameters in text form")
	common.AddRandomStateFlags(ecparamCmd)
	common.AddProviderFlags(ecparamCmd)
	rootCmd.AddCommand(ecparamCmd)

	carapace.Gen(ecparamCmd).FlagCompletion(carapace.ActionMap{
		"conv_form": carapace.ActionValues("compressed", "hybrid", "uncompressed"),
		"in":        carapace.ActionFiles(),
		"inform":    carapace.ActionValues("DER", "PEM"),
		"out":       carapace.ActionFiles(),
		"outform":   carapace.ActionValues("DER", "PEM"),
		"param_enc": carapace.ActionValues("explicit", "named_curve"),
	})
}
