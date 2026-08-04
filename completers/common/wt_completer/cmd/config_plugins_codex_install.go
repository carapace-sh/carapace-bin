package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_codex_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Configure the Worktrunk marketplace in Codex",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_codex_installCmd).Standalone()

	config_plugins_codex_installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_plugins_codexCmd.AddCommand(config_plugins_codex_installCmd)
}
