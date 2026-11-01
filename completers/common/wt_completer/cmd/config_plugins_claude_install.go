package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_claude_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the Worktrunk plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_claude_installCmd).Standalone()

	config_plugins_claude_installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_plugins_claudeCmd.AddCommand(config_plugins_claude_installCmd)
}
