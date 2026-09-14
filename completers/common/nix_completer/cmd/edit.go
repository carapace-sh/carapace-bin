package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "open the Nix expression of a Nix package in $EDITOR",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	rootCmd.AddCommand(editCmd)

	common.AddEvaluationFlags(editCmd)
	common.AddFlakeFlags(editCmd)
	common.AddInterpretationFlags(editCmd)
	common.AddLoggingFlags(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(editCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
