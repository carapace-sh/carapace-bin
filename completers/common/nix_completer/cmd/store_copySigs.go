package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_copySigsCmd = &cobra.Command{
	Use:   "copy-sigs",
	Short: "copy store path signatures from substituters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_copySigsCmd).Standalone()

	store_copySigsCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	store_copySigsCmd.Flags().StringP("substituter", "s", "", "Copy signatures from the specified store")
	storeCmd.AddCommand(store_copySigsCmd)

	addEvaluationFlags(store_copySigsCmd)
	addFlakeFlags(store_copySigsCmd)
	addLoggingFlags(store_copySigsCmd)

	carapace.Gen(store_copySigsCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(store_copySigsCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
