package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_copyLogCmd = &cobra.Command{
	Use:   "copy-log",
	Short: "copy build logs between Nix stores",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_copyLogCmd).Standalone()

	store_copyLogCmd.Flags().String("from", "", "URL of the source Nix store")
	store_copyLogCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	store_copyLogCmd.Flags().String("to", "", "URL of the destination Nix store")
	storeCmd.AddCommand(store_copyLogCmd)

	addEvaluationFlags(store_copyLogCmd)
	addFlakeFlags(store_copyLogCmd)
	addLoggingFlags(store_copyLogCmd)

	carapace.Gen(store_copyLogCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(store_copyLogCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
