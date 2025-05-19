package cmd

import (
	"github.com/carapace-sh/carapace"
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

	printDevEnvCmd.Flags().Bool("json", false, "Produce output in JSON format")
	printDevEnvCmd.Flags().String("profile", "", "The profile to operate on")
	printDevEnvCmd.Flags().String("redirect", "", "Redirect a store path to a mutable location")
	rootCmd.AddCommand(printDevEnvCmd)

	addEvaluationFlags(printDevEnvCmd)
	addFlakeFlags(printDevEnvCmd)
	addLoggingFlags(printDevEnvCmd)

	carapace.Gen(printDevEnvCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"profile":             carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(printDevEnvCmd).PositionalCompletion(nix.ActionInstallables())
}
