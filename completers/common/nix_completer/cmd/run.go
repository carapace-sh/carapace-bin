package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
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

	runCmd.Flags().BoolP("ignore-env", "i", false, "Clear the entire environment, except for those specified with --keep-env-var")
	runCmd.Flags().StringP("keep-env-var", "k", "", "Keep the environment variable name, when using --ignore-env")
	runCmd.Flags().String("set-env-var", "", "Sets an environment variable name with value")
	runCmd.Flags().StringP("unset-env-var", "u", "", "Unset the environment variable name")
	runCmd.Flag("set-env-var").Nargs = 2

	common.AddEvaluationFlags(runCmd)
	common.AddFlakeFlags(runCmd)
	common.AddInterpretationFlags(runCmd)
	common.AddLoggingFlags(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"keep-env-var":  os.ActionEnvironmentVariables(),
		"set-env-var":   action.ActionSetEnvVar(),
		"unset-env-var": os.ActionEnvironmentVariables(),
	})
	carapace.Gen(runCmd).PositionalCompletion(nix.ActionInstallables())
}
