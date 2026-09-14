package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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

	derivation_showCmd.Flags().BoolP("recursive", "r", false, "Include the dependencies of the specified derivations")
	derivation_showCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	derivationCmd.AddCommand(derivation_showCmd)

	common.AddEvaluationFlags(derivation_showCmd)
	common.AddFlakeFlags(derivation_showCmd)
	common.AddInterpretationFlags(derivation_showCmd)
	common.AddLoggingFlags(derivation_showCmd)

	carapace.Gen(derivation_showCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(derivation_showCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
