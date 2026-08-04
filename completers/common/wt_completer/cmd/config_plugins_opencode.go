package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_opencodeCmd = &cobra.Command{
	Use:   "opencode",
	Short: "OpenCode plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_opencodeCmd).Standalone()

	config_plugins_opencodeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pluginsCmd.AddCommand(config_plugins_opencodeCmd)
}
