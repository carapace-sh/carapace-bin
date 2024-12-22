package cmd

import (
	"github.com/carapace-sh/carapace"
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
	shellCmd.Flags().BoolP("ignore-environment", "i", false, "Clear the entire environment")
	shellCmd.Flags().StringP("keep", "k", "", "Keep the environment variable name")
	shellCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	shellCmd.Flags().StringP("unset", "u", "", "Unset the environment variable name")
	rootCmd.AddCommand(shellCmd)

	shellCmd.Flag("command").Nargs = -1

	addEvaluationFlags(shellCmd)
	addFlakeFlags(shellCmd)
	addLoggingFlags(shellCmd)

	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"keep":                os.ActionEnvironmentVariables(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
		"unset":               os.ActionEnvironmentVariables(),
	})
	carapace.Gen(shellCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
