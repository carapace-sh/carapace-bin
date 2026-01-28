package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dsaparamCmd = &cobra.Command{
	Use:     "dsaparam",
	Short:   "DSA Parameter Generation and Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsaparamCmd).Standalone()

	dsaparamCmd.Flags().BoolS("genkey", "genkey", false, "Generate a DSA key")
	dsaparamCmd.Flags().StringS("in", "in", "", "Input file")
	dsaparamCmd.Flags().StringS("inform", "inform", "", "Input format - DER or PEM")
	dsaparamCmd.Flags().BoolS("noout", "noout", false, "No output")
	dsaparamCmd.Flags().StringS("out", "out", "", "Output file")
	dsaparamCmd.Flags().StringS("outform", "outform", "", "Output format - DER or PEM")
	dsaparamCmd.Flags().BoolS("quiet", "quiet", false, "Terse output")
	dsaparamCmd.Flags().BoolS("text", "text", false, "Print as text")
	dsaparamCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	common.AddProviderFlags(dsaparamCmd)
	common.AddRandomStateFlags(dsaparamCmd)
	rootCmd.AddCommand(dsaparamCmd)

	carapace.Gen(dsaparamCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
