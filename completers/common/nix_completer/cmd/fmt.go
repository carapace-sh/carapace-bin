package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:     "fmt",
	Short:   "reformat your code in the standard style",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	rootCmd.AddCommand(fmtCmd)

	common.AddEvaluationFlags(fmtCmd)
	common.AddFlakeFlags(fmtCmd)
	common.AddInterpretationFlags(fmtCmd)
	common.AddLoggingFlags(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(fmtCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
