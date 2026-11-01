package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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

	common.AddEvaluationFlags(logCmd)
	common.AddFlakeFlags(logCmd)
	common.AddInterpretationFlags(logCmd)
	common.AddLoggingFlags(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(logCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
