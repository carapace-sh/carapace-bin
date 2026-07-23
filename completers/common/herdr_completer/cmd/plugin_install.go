package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install <source>",
	Short: "Install a plugin from GitHub",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()

	plugin_installCmd.Flags().String("ref", "", "")
	plugin_installCmd.Flags().BoolP("yes", "y", false, "")
	pluginCmd.AddCommand(plugin_installCmd)
}
