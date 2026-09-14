package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_makeContentAddressedCmd = &cobra.Command{
	Use:   "make-content-addressed",
	Short: "rewrite a path or closure to content-addressed form",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_makeContentAddressedCmd).Standalone()

	store_makeContentAddressedCmd.Flags().String("from", "", "URL of the source Nix store")
	store_makeContentAddressedCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	store_makeContentAddressedCmd.Flags().String("to", "", "URL of the destination Nix store")
	storeCmd.AddCommand(store_makeContentAddressedCmd)

	common.AddBuiltPathsFlags(store_makeContentAddressedCmd)
	common.AddEvaluationFlags(store_makeContentAddressedCmd)
	common.AddFlakeFlags(store_makeContentAddressedCmd)
	common.AddLoggingFlags(store_makeContentAddressedCmd)

	// TODO: --from/--to flag completion
	carapace.Gen(store_makeContentAddressedCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_makeContentAddressedCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
