package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shell_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Write shell integration to config files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shell_installCmd).Standalone()

	config_shell_installCmd.Flags().String("cmd", "", "Command name for shell integration (defaults to binary name)")
	config_shell_installCmd.Flags().Bool("dry-run", false, "Show what would be changed")
	config_shell_installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_shell_installCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	config_shellCmd.AddCommand(config_shell_installCmd)
}
