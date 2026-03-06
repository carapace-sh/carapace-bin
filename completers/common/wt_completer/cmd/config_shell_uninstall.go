package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shell_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove shell integration from config files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shell_uninstallCmd).Standalone()

	config_shell_uninstallCmd.Flags().Bool("dry-run", false, "Show what would be changed")
	config_shell_uninstallCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_shell_uninstallCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	config_shellCmd.AddCommand(config_shell_uninstallCmd)

	carapace.Gen(config_shell_uninstallCmd).PositionalCompletion(
		carapace.ActionValues("bash", "fish", "nu", "zsh", "powershell"),
	)
}
