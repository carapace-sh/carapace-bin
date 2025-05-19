package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_dumpPathCmd = &cobra.Command{
	Use:   "dump-path",
	Short: "serialise a store path to stdout in NAR format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_dumpPathCmd).Standalone()

	store_dumpPathCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	storeCmd.AddCommand(store_dumpPathCmd)

	addEvaluationFlags(store_dumpPathCmd)
	addFlakeFlags(store_dumpPathCmd)
	addLoggingFlags(store_dumpPathCmd)

	carapace.Gen(store_dumpPathCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(store_dumpPathCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
