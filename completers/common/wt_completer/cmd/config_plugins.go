package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Plugin management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pluginsCmd).Standalone()

	config_pluginsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_pluginsCmd)
}
