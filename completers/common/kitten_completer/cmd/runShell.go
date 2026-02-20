package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runShellCmd = &cobra.Command{
	Use:   "run-shell",
	Short: "Run the user's shell with shell integration enabled",
}

func init() {
	rootCmd.AddCommand(runShellCmd)
	carapace.Gen(runShellCmd).Standalone()

	runShellCmd.Flags().String("shell-integration", "", "Specify a value for the shell_integration option, overriding the one from kitty.conf.")
	runShellCmd.Flags().String("shell", "", "Specify the shell command to run. The default value of . will use the parent shell if recognized, falling back to the value of the shell option from kitty.conf.")
	runShellCmd.Flags().StringArray("env", nil, "Specify an env var to set before running the shell. Of the form KEY=VAL. Can be specified multiple times. If no = is present KEY is unset.")
	runShellCmd.Flags().String("cwd", "", "The working directory to use when executing the shell.")
	runShellCmd.Flags().String("inject-self-onto-path", "", "Add the directory containing this kitten binary to PATH. Directory is added only if not already present.")
	runShellCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(runShellCmd).FlagCompletion(carapace.ActionMap{
		"inject-self-onto-path": carapace.ActionValues("always", "never", "unless-root"),
	})

	carapace.Gen(runShellCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
