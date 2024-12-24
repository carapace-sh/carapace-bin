package cmd

import (
	"github.com/carapace-sh/carapace"
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

	addEvaluationFlags(replCmd)
	addFlakeFlags(replCmd)
	addInterpretationFlags(replCmd)
	addLoggingFlags(replCmd)

	carapace.Gen(replCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(replCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
