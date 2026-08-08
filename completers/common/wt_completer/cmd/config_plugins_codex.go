package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_codexCmd = &cobra.Command{
	Use:   "codex",
	Short: "Codex plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_codexCmd).Standalone()

	config_plugins_codexCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pluginsCmd.AddCommand(config_plugins_codexCmd)
}
