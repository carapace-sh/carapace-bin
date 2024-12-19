package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/nix_completer/cmd/action"
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
	developCmd.Flags().BoolP("ignore-environment", "i", false, "Clear the entire environment (except those specified wi--keep).    ")
	developCmd.Flags().Bool("install", false, "Run the install phase.")
	developCmd.Flags().Bool("installcheck", false, "Run the installcheck phase.")
	developCmd.Flags().StringP("keep", "k", "", "Keep the environment variable name.")
	developCmd.Flags().String("phase", "", "The stdenv phase to run.")
	developCmd.Flags().String("profile", "", "The profile to update.")
	developCmd.Flags().String("redirect", "", "Redirect a store path to a mutable location.")
	developCmd.Flags().Bool("unpack", false, "Run the unpack phase.")
	developCmd.Flags().StringP("unset", "u", "", "Unset the environment variable name.")

	addEvaluationFlags(developCmd)
	addFlakeFlags(developCmd)
	addInterpretationFlags(developCmd)
	addLoggingFlags(developCmd)

	rootCmd.AddCommand(developCmd)

	carapace.Gen(developCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"inputs-from":         nix.ActionFlakes(),
		"keep":                os.ActionEnvironmentVariables(),
		"output-lock-file":    carapace.ActionFiles(),
		"profile":             carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
		"unset":               os.ActionEnvironmentVariables(),
	})
	carapace.Gen(developCmd).PositionalCompletion(action.ActionFlakeRefs([]string{"nix", "develop"}))
}
