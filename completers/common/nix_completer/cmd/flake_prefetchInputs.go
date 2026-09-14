package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_prefetchInputsCmd = &cobra.Command{
	Use:   "prefetch-inputs [flags] [flake-url]",
	Short: "fetch the inputs of a flake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_prefetchInputsCmd).Standalone()

	common.AddEvaluationFlags(flake_prefetchInputsCmd)
	common.AddFlakeFlags(flake_prefetchInputsCmd)
	common.AddLoggingFlags(flake_prefetchInputsCmd)

	flakeCmd.AddCommand(flake_prefetchInputsCmd)

	carapace.Gen(flake_prefetchInputsCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(flake_prefetchInputsCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())
}
