package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var derivation_showCmd = &cobra.Command{
	Use:   "show",
	Short: "work with derivations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(derivation_showCmd).Standalone()

	derivation_showCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	derivation_showCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	derivationCmd.AddCommand(derivation_showCmd)

	addEvaluationFlags(derivation_showCmd)
	addFlakeFlags(derivation_showCmd)
	addLoggingFlags(derivation_showCmd)

	carapace.Gen(derivation_showCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(derivation_showCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
