package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_metadataCmd = &cobra.Command{
	Use:     "metadata [flags] [flake-url]",
	Short:   "show flake metadata",
	Run:     func(cmd *cobra.Command, args []string) {},
	Aliases: []string{"info"},
}

func init() {
	carapace.Gen(flake_metadataCmd).Standalone()

	common.AddEvaluationFlags(flake_metadataCmd)
	common.AddFlakeFlags(flake_metadataCmd)
	common.AddLoggingFlags(flake_metadataCmd)

	carapace.Gen(flake_metadataCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(flake_metadataCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_metadataCmd)
}
