package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run a Nix application",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	rootCmd.AddCommand(runCmd)

	addEvaluationFlags(runCmd)
	addFlakeFlags(runCmd)
	addLoggingFlags(runCmd)

	// TODO positional completion
}
