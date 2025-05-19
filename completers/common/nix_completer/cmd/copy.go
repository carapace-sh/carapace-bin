package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Short:   "copy paths between Nix stores",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyCmd).Standalone()

	copyCmd.Flags().String("from", "", "URL of the source Nix store")
	copyCmd.Flags().Bool("no-check-sigs", false, "Do not require that paths are signed by trusted keys")
	copyCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	copyCmd.Flags().BoolP("substitute-on-destination", "s", false, "Whether to try substitutes on the destination store")
	copyCmd.Flags().String("to", "", "URL of the destination Nix store")
	rootCmd.AddCommand(copyCmd)

	addEvaluationFlags(copyCmd)
	addFlakeFlags(copyCmd)
	addLoggingFlags(copyCmd)

	// TODO: --from/--to completion
	carapace.Gen(copyCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})

	carapace.Gen(copyCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
