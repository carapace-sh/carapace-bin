package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_showCmd = &cobra.Command{
	Use:   "show [flags] [flake-url]",
	Short: "show the outputs provided by a flake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_showCmd).Standalone()

	flake_showCmd.Flags().Bool("all-systems", false, "Show the contents of outputs for all systems")
	flake_showCmd.Flags().Bool("json", false, "Produce output in JSON format")
	flake_showCmd.Flags().Bool("legacy", false, "Show the contents of the legacyPackages output")

	addEvaluationFlags(flake_showCmd)
	addFlakeFlags(flake_showCmd)
	addLoggingFlags(flake_showCmd)

	carapace.Gen(flake_showCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(flake_showCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_showCmd)
}
