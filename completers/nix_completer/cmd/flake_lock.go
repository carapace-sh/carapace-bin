package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_lockCmd = &cobra.Command{
	Use:   "lock [flags] [flake-url]",
	Short: "create missing lockfile entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_lockCmd).Standalone()

	addEvaluationFlags(flake_lockCmd)
	addFlakeFlags(flake_lockCmd)
	addLoggingFlags(flake_lockCmd)

	carapace.Gen(flake_lockCmd).PositionalCompletion(nix.ActionFlakes())

	flakeCmd.AddCommand(flake_lockCmd)
}
