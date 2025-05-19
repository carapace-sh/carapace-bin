package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_checkCmd = &cobra.Command{
	Use:   "check [flags] [flake-url]",
	Short: "check whether the flake evaluates and runs its tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_checkCmd).Standalone()

	flake_checkCmd.Flags().Bool("all-systems", false, "Check the outputs for all systems")
	flake_checkCmd.Flags().Bool("no-build", false, "Do not build checks")

	addEvaluationFlags(flake_checkCmd)
	addFlakeFlags(flake_checkCmd)
	addLoggingFlags(flake_checkCmd)

	carapace.Gen(flake_checkCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(flake_checkCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_checkCmd)
}
