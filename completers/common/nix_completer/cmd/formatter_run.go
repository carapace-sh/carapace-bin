package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var formatter_runCmd = &cobra.Command{
	Use:   "run [flags] [installable]",
	Short: "run the default build of the flake attribute formatter.<system>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatter_runCmd).Standalone()

	common.AddJSONFlags(formatter_runCmd)
	common.AddEvaluationFlags(formatter_runCmd)
	common.AddFlakeFlags(formatter_runCmd)
	common.AddInterpretationFlags(formatter_runCmd)
	common.AddLoggingFlags(formatter_runCmd)

	formatterCmd.AddCommand(formatter_runCmd)

	carapace.Gen(formatter_runCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(formatter_runCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
