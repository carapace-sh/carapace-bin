package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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
	flake_showCmd.Flags().Bool("legacy", false, "Show the contents of the legacyPackages output")

	common.AddEvaluationFlags(flake_showCmd)
	common.AddFlakeFlags(flake_showCmd)
	common.AddLoggingFlags(flake_showCmd)

	carapace.Gen(flake_showCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(flake_showCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_showCmd)
}
