package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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

	store_diffClosuresCmd.Flags().Bool("derivation", false, "Operate on the store derivation rather than its outputs")
	storeCmd.AddCommand(store_diffClosuresCmd)

	common.AddEvaluationFlags(store_diffClosuresCmd)
	common.AddFlakeFlags(store_diffClosuresCmd)
	common.AddInterpretationFlags(store_diffClosuresCmd)
	common.AddLoggingFlags(store_diffClosuresCmd)

	carapace.Gen(store_diffClosuresCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_diffClosuresCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
