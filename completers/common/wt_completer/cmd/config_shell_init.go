package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shell_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate shell integration code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shell_initCmd).Standalone()

	config_shell_initCmd.Flags().String("cmd", "", "Command name for shell integration (defaults to binary name)")
	config_shell_initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_shellCmd.AddCommand(config_shell_initCmd)

	carapace.Gen(config_shell_initCmd).FlagCompletion(carapace.ActionMap{
		"cmd": carapace.ActionExecutables(), // TODO what is expected here?
	})

	carapace.Gen(config_shell_initCmd).PositionalCompletion(
		carapace.ActionValues("bash", "fish", "nu", "zsh", "powershell"),
	)
}
