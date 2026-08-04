package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_claude_installStatuslineCmd = &cobra.Command{
	Use:   "install-statusline",
	Short: "Configure the Claude Code statusline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_claude_installStatuslineCmd).Standalone()

	config_plugins_claude_installStatuslineCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_plugins_claudeCmd.AddCommand(config_plugins_claude_installStatuslineCmd)
}
