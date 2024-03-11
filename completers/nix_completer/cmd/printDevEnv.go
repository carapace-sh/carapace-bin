package cmd

import (
	"github.com/carapace-sh/carapace"
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
	rootCmd.AddCommand(printDevEnvCmd)

	addEvaluationFlags(printDevEnvCmd)
	addFlakeFlags(printDevEnvCmd)
	addLoggingFlags(printDevEnvCmd)

	carapace.Gen(printDevEnvCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionFiles(),
	})

	// TODO positional completion
}
