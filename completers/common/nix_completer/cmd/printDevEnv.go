package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var printDevEnvCmd = &cobra.Command{
	Use:     "print-dev-env",
	Short:   "print shell code that can be sourced by bash to reproduce the build environment of a derivation",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printDevEnvCmd).Standalone()

	printDevEnvCmd.Flags().String("profile", "", "The profile to operate on")
	printDevEnvCmd.Flags().String("redirect", "", "Redirect a store path to a mutable location")
	rootCmd.AddCommand(printDevEnvCmd)

	common.AddEvaluationFlags(printDevEnvCmd)
	common.AddFlakeFlags(printDevEnvCmd)
	common.AddLoggingFlags(printDevEnvCmd)

	carapace.Gen(printDevEnvCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionFiles(),
	})
	carapace.Gen(printDevEnvCmd).PositionalCompletion(nix.ActionInstallables())
}
