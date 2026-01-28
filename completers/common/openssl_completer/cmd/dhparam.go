package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dhparamCmd = &cobra.Command{
	Use:     "dhparam",
	Short:   "Generation and Management of Diffie-Hellman Parameters",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dhparamCmd).Standalone()

	dhparamCmd.Flags().BoolS("2", "2", false, "Generate parameters using 2 as the generator value")
	dhparamCmd.Flags().BoolS("3", "3", false, "Generate parameters using 3 as the generator value")
	dhparamCmd.Flags().BoolS("5", "5", false, "Generate parameters using 5 as the generator value")
	dhparamCmd.Flags().BoolS("check", "check", false, "Check the DH parameters")
	dhparamCmd.Flags().BoolS("dsaparam", "dsaparam", false, "Read or generate DSA parameters, convert to DH")
	dhparamCmd.Flags().StringS("in", "in", "", "Input file")
	dhparamCmd.Flags().StringS("inform", "inform", "", "Input format, DER or PEM")
	dhparamCmd.Flags().BoolS("noout", "noout", false, "Don't output any DH parameters")
	dhparamCmd.Flags().StringS("out", "out", "", "Output file")
	dhparamCmd.Flags().StringS("outform", "outform", "", "Output format, DER or PEM")
	dhparamCmd.Flags().BoolS("quiet", "quiet", false, "Terse output")
	dhparamCmd.Flags().BoolS("text", "text", false, "Print a text form of the DH parameters")
	dhparamCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	common.AddProviderFlags(dhparamCmd)
	common.AddRandomStateFlags(dhparamCmd)
	rootCmd.AddCommand(dhparamCmd)

	carapace.Gen(dhparamCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
