package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_updateCmd = &cobra.Command{
	Use:   "update [flags] [flake-url]",
	Short: "update flake lock file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_updateCmd).Standalone()

	addEvaluationFlags(flake_updateCmd)
	addFlakeFlags(flake_updateCmd)
	addLoggingFlags(flake_updateCmd)

	carapace.Gen(flake_updateCmd).PositionalCompletion(nix.ActionFlakes())

	flakeCmd.AddCommand(flake_updateCmd)
}
