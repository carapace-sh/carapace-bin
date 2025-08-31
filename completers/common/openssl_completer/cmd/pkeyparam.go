package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkeyparamCmd = &cobra.Command{
	Use:     "pkeyparam",
	Short:   "Public key algorithm parameter management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkeyparamCmd).Standalone()

	pkeyparamCmd.Flags().BoolS("check", "check", false, "Check key param consistency")
	pkeyparamCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	pkeyparamCmd.Flags().StringS("in", "in", "", "Input file")
	pkeyparamCmd.Flags().BoolS("noout", "noout", false, "Don't output encoded parameters")
	pkeyparamCmd.Flags().StringS("out", "out", "", "Output file")
	pkeyparamCmd.Flags().BoolS("text", "text", false, "Print parameters as text")
	common.AddProviderFlags(pkeyparamCmd)
	rootCmd.AddCommand(pkeyparamCmd)

	carapace.Gen(pkeyparamCmd).FlagCompletion(carapace.ActionMap{
		"engine": action.ActionEngines(),
		"in":     carapace.ActionFiles(),
		"out":    carapace.ActionFiles(),
	})
}
