package cmd

import (
	"github.com/carapace-sh/carapace"
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

	store_signCmd.Flags().BoolP("key-file", "k", false, "File containing the secret signing key")
	store_signCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	storeCmd.AddCommand(store_signCmd)

	addEvaluationFlags(store_signCmd)
	addFlakeFlags(store_signCmd)
	addLoggingFlags(store_signCmd)

	carapace.Gen(store_signCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"key-file":            carapace.ActionFiles(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})

	carapace.Gen(store_signCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
