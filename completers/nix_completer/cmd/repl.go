package cmd

import (
	"github.com/carapace-sh/carapace"
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

	addEvaluationFlags(replCmd)
	addFlakeFlags(replCmd)
	addInterpretationFlags(replCmd)
	addLoggingFlags(replCmd)

	// TODO positional completion
}
