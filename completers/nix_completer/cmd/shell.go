package cmd

import (
	"github.com/rsteube/carapace"
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
	shellCmd.Flags().StringP("unset", "u", "", "Unset the environment variable name")
	rootCmd.AddCommand(shellCmd)

	shellCmd.Flag("command").Nargs = -1

	addEvaluationFlags(shellCmd)
	addFlakeFlags(shellCmd)
	addLoggingFlags(shellCmd)

	// TODO flag completion

	// TODO positional completion
}
