package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:     "repl",
	Short:   "start an interactive environment for evaluating Nix expressions",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replCmd).Standalone()

	rootCmd.AddCommand(replCmd)

	replCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	common.AddEvaluationFlags(replCmd)
	common.AddFlakeFlags(replCmd)
	common.AddInterpretationFlags(replCmd)
	common.AddLoggingFlags(replCmd)

	carapace.Gen(replCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(replCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
