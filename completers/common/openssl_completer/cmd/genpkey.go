package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var genpkeyCmd = &cobra.Command{
	Use:     "genpkey",
	Short:   "Generation of Private Key or Parameters",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genpkeyCmd).Standalone()

	genpkeyCmd.Flags().StringS("algorithm", "algorithm", "", "The public key algorithm")
	genpkeyCmd.Flags().StringS("config", "config", "", "Load a configuration file (this may load modules)")
	genpkeyCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	genpkeyCmd.Flags().BoolS("genparam", "genparam", false, "Generate parameters, not key")
	genpkeyCmd.Flags().StringS("out", "out", "", "Output (private key) file")
	genpkeyCmd.Flags().StringS("outform", "outform", "", "output format (DER or PEM)")
	genpkeyCmd.Flags().StringS("outpubkey", "outpubkey", "", "Output public key file")
	genpkeyCmd.Flags().StringS("paramfile", "paramfile", "", "Parameters file")
	genpkeyCmd.Flags().StringS("pass", "pass", "", "Output file pass phrase source")
	genpkeyCmd.Flags().StringSliceS("pkeyopt", "pkeyopt", nil, "Set the public key algorithm option as opt:value")
	genpkeyCmd.Flags().BoolS("quiet", "quiet", false, "Do not output status while generating keys")
	genpkeyCmd.Flags().BoolS("text", "text", false, "Print the private key in text")
	genpkeyCmd.Flags().BoolS("verbose", "verbose", false, "Output status while generating keys")
	common.AddProviderFlags(genpkeyCmd)
	common.AddRandomStateFlags(genpkeyCmd)
	rootCmd.AddCommand(genpkeyCmd)

	genpkeyCmd.MarkFlagsMutuallyExclusive("algorithm", "paramfile")

	carapace.Gen(genpkeyCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(),
		"engine":    action.ActionEngines(),
		"out":       carapace.ActionFiles(),
		"outform":   carapace.ActionValues("DER", "PEM"),
		"outpubkey": carapace.ActionFiles(),
		"paramfile": carapace.ActionFiles(),
	})
}
