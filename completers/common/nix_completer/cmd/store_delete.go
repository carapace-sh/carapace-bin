package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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

	store_deleteCmd.Flags().Bool("also-referrers", false, "Also allow deletion of any referrers of the specified paths")
	store_deleteCmd.Flags().Bool("ignore-liveness", false, "Do not check whether the paths are reachable from a root")
	store_deleteCmd.Flags().Bool("skip-alive", false, "Do not emit errors when attempting to delete something that is still alive")
	store_deleteCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	storeCmd.AddCommand(store_deleteCmd)

	common.AddBuiltPathsFlags(store_deleteCmd)
	common.AddEvaluationFlags(store_deleteCmd)
	common.AddFlakeFlags(store_deleteCmd)
	common.AddLoggingFlags(store_deleteCmd)

	carapace.Gen(store_deleteCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_deleteCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
