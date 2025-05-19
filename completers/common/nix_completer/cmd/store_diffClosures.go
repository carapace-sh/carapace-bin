package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_diffClosuresCmd = &cobra.Command{
	Use:   "diff-closures",
	Short: "show what packages and versions were added and removed between two closures",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_diffClosuresCmd).Standalone()

	storeCmd.AddCommand(store_diffClosuresCmd)

	addEvaluationFlags(store_diffClosuresCmd)
	addFlakeFlags(store_diffClosuresCmd)
	addLoggingFlags(store_diffClosuresCmd)

	carapace.Gen(store_diffClosuresCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	// TODO positional completion
}
