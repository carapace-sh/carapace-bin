package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_claudeCmd = &cobra.Command{
	Use:   "claude",
	Short: "Claude Code plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_claudeCmd).Standalone()

	config_plugins_claudeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pluginsCmd.AddCommand(config_plugins_claudeCmd)
}
