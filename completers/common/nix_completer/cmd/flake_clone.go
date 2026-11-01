package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_cloneCmd = &cobra.Command{
	Use:   "clone [flags] [flake-url]",
	Short: "clone flake repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_cloneCmd).Standalone()

	flake_cloneCmd.Flags().StringP("dest", "f", "", "Specify target directory for flake repository")

	common.AddEvaluationFlags(flake_cloneCmd)
	common.AddFlakeFlags(flake_cloneCmd)
	common.AddLoggingFlags(flake_cloneCmd)

	carapace.Gen(flake_cloneCmd).FlagCompletion(carapace.ActionMap{
		"dest": carapace.ActionDirectories(),
	})
	carapace.Gen(flake_cloneCmd).PositionalCompletion(carapace.Batch(
		carapace.ActionDirectories(),
		nix.ActionFlakes(),
	).ToA())

	flakeCmd.AddCommand(flake_cloneCmd)
}
