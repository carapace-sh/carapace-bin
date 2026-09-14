package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "run a shell in which the specified packages are available",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().StringP("command", "c", "", "Command and arguments to be executed, defaulting to $SHELL")
	shellCmd.Flags().BoolP("ignore-env", "i", false, "Clear the entire environment, except for those specified with --keep-env-var")
	shellCmd.Flags().StringP("keep-env-var", "k", "", "Keep the environment variable name, when using --ignore-env")
	shellCmd.Flags().String("set-env-var", "", "Sets an environment variable name with value")
	shellCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	shellCmd.Flags().StringP("unset-env-var", "u", "", "Unset the environment variable name")
	rootCmd.AddCommand(shellCmd)

	shellCmd.Flag("command").Nargs = -1
	shellCmd.Flag("set-env-var").Nargs = 2

	common.AddEvaluationFlags(shellCmd)
	common.AddFlakeFlags(shellCmd)
	common.AddInterpretationFlags(shellCmd)
	common.AddLoggingFlags(shellCmd)

	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"keep-env-var": os.ActionEnvironmentVariables(),
		"set-env-var":  action.ActionSetEnvVar(),
	})
	carapace.Gen(shellCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
