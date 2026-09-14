package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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

	common.AddBuiltPathsFlags(store_dumpPathCmd)
	common.AddEvaluationFlags(store_dumpPathCmd)
	common.AddFlakeFlags(store_dumpPathCmd)
	common.AddLoggingFlags(store_dumpPathCmd)

	carapace.Gen(store_dumpPathCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_dumpPathCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
