package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_codex_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove the Worktrunk marketplace from Codex",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_codex_uninstallCmd).Standalone()

	config_plugins_codex_uninstallCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_plugins_codexCmd.AddCommand(config_plugins_codex_uninstallCmd)
}
