package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var env_shellCmd = &cobra.Command{
	Use:   "shell [flags] [installable]",
	Short: "run a shell in which the specified packages are available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_shellCmd).Standalone()

	env_shellCmd.Flags().StringP("command", "c", "", "Command and arguments to be executed, defaulting to $SHELL")
	env_shellCmd.Flags().BoolP("ignore-env", "i", false, "Clear the entire environment, except for those specified with --keep-env-var")
	env_shellCmd.Flags().StringP("keep-env-var", "k", "", "Keep the environment variable name, when using --ignore-env")
	env_shellCmd.Flags().String("set-env-var", "", "Sets an environment variable name with value")
	env_shellCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	env_shellCmd.Flags().StringP("unset-env-var", "u", "", "Unset the environment variable name")
	envCmd.AddCommand(env_shellCmd)

	env_shellCmd.Flag("command").Nargs = -1
	env_shellCmd.Flag("set-env-var").Nargs = 2

	common.AddEvaluationFlags(env_shellCmd)
	common.AddFlakeFlags(env_shellCmd)
	common.AddInterpretationFlags(env_shellCmd)
	common.AddLoggingFlags(env_shellCmd)

	carapace.Gen(env_shellCmd).FlagCompletion(carapace.ActionMap{
		"keep-env-var": os.ActionEnvironmentVariables(),
		"set-env-var":  action.ActionSetEnvVar(),
	})
	carapace.Gen(env_shellCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
