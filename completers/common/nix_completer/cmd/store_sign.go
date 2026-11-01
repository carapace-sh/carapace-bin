package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_signCmd = &cobra.Command{
	Use:   "sign",
	Short: "sign store paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_signCmd).Standalone()

	store_signCmd.Flags().StringP("key-file", "k", "", "File containing the secret signing key")
	store_signCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	storeCmd.AddCommand(store_signCmd)

	common.AddBuiltPathsFlags(store_signCmd)
	common.AddEvaluationFlags(store_signCmd)
	common.AddFlakeFlags(store_signCmd)
	common.AddLoggingFlags(store_signCmd)

	carapace.Gen(store_signCmd).FlagCompletion(carapace.ActionMap{
		"key-file": carapace.ActionFiles(),
	})

	carapace.Gen(store_signCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
