package cmd

import (
	"github.com/carapace-sh/carapace"
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

	addEvaluationFlags(store_repairCmd)
	addFlakeFlags(store_repairCmd)
	addLoggingFlags(store_repairCmd)

	carapace.Gen(store_repairCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(store_repairCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
