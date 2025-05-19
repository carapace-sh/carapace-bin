package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete paths from the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_deleteCmd).Standalone()

	store_deleteCmd.Flags().Bool("ignore-liveness", false, "Do not check whether the paths are reachable from a root")
	store_deleteCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	storeCmd.AddCommand(store_deleteCmd)

	addEvaluationFlags(store_deleteCmd)
	addFlakeFlags(store_deleteCmd)
	addLoggingFlags(store_deleteCmd)

	carapace.Gen(store_deleteCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(store_deleteCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
