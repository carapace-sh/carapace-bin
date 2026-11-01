package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var developCmd = &cobra.Command{
	Use:     "develop",
	Short:   "run a bash shell that provides the build environment of a derivation",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developCmd).Standalone()

	developCmd.Flags().Bool("build", false, "Run the build phase.")
	developCmd.Flags().Bool("check", false, "Run the check phase.")
	developCmd.Flags().StringP("command", "c", "", "Instead of starting an interactive shell, start the specified command and arguments.")
	developCmd.Flags().Bool("configure", false, "Run the configure phase.")
	developCmd.Flags().BoolP("ignore-env", "i", false, "Clear the entire environment, except for those specified with --keep-env-var")
	developCmd.Flags().Bool("install", false, "Run the install phase.")
	developCmd.Flags().Bool("installcheck", false, "Run the installcheck phase.")
	developCmd.Flags().StringP("keep-env-var", "k", "", "Keep the environment variable name, when using --ignore-env")
	developCmd.Flags().String("phase", "", "The stdenv phase to run.")
	developCmd.Flags().String("profile", "", "The profile to update.")
	developCmd.Flags().String("redirect", "", "Redirect a store path to a mutable location.")
	developCmd.Flags().String("set-env-var", "", "Sets an environment variable name with value")
	developCmd.Flags().Bool("unpack", false, "Run the unpack phase.")
	developCmd.Flags().StringP("unset-env-var", "u", "", "Unset the environment variable name")

	developCmd.Flag("set-env-var").Nargs = 2

	common.AddEvaluationFlags(developCmd)
	common.AddFlakeFlags(developCmd)
	common.AddInterpretationFlags(developCmd)
	common.AddLoggingFlags(developCmd)

	rootCmd.AddCommand(developCmd)

	carapace.Gen(developCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"keep-env-var": os.ActionEnvironmentVariables(),
		"profile":      carapace.ActionFiles(),
		"set-env-var":  action.ActionSetEnvVar(),
	})

	carapace.Gen(developCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakeRefs(),
		).ToA(),
	)
}
