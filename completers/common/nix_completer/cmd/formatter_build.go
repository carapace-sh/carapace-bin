package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var formatter_buildCmd = &cobra.Command{
	Use:   "build [flags] [installable]",
	Short: "build the flake attribute formatter.<system> and print the formatter program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatter_buildCmd).Standalone()

	common.AddEvaluationFlags(formatter_buildCmd)
	common.AddFlakeFlags(formatter_buildCmd)
	common.AddInterpretationFlags(formatter_buildCmd)
	common.AddLoggingFlags(formatter_buildCmd)

	formatterCmd.AddCommand(formatter_buildCmd)

	carapace.Gen(formatter_buildCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(formatter_buildCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
