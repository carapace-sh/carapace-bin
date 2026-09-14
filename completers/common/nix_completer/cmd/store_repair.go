package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "repair store path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_repairCmd).Standalone()

	store_repairCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	storeCmd.AddCommand(store_repairCmd)

	common.AddBuiltPathsFlags(store_repairCmd)
	common.AddEvaluationFlags(store_repairCmd)
	common.AddFlakeFlags(store_repairCmd)
	common.AddLoggingFlags(store_repairCmd)

	carapace.Gen(store_repairCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_repairCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
