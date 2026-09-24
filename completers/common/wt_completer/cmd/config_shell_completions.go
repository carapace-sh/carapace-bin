package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shell_completionsCmd = &cobra.Command{
	Use:    "completions",
	Short:  "Generate shell completions for package managers",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shell_completionsCmd).Standalone()

	config_shell_completionsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_shellCmd.AddCommand(config_shell_completionsCmd)

	carapace.Gen(config_shell_completionsCmd).PositionalCompletion(
		carapace.ActionValues("bash", "fish", "nu", "zsh", "powershell"),
	)
}
