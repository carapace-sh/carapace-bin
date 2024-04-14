package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:     "log",
	Short:   "show the build log of the specified packages or paths, if available",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	rootCmd.AddCommand(logCmd)

	addEvaluationFlags(logCmd)
	addFlakeFlags(logCmd)
	addLoggingFlags(logCmd)

	carapace.Gen(logCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
