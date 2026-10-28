package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_getCmd).Standalone()

	plugin_getCmd.Flags().String("arch", "", "The architecture to pull the plugin for, give a value or leave blank for auto-detection")
	plugin_getCmd.Flags().String("os", "", "The OS to pull the plugin for, give a value or leave blank for auto-detection")
	plugin_getCmd.Flags().String("path", "$HOME/.openfaas/plugins", "The path for the plugin")
	plugin_getCmd.Flags().String("registry", "ghcr.io/openfaasltd", "The registry to pull the plugin from")
	plugin_getCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	plugin_getCmd.Flags().String("version", "latest", "Version or SHA for plugin")
	pluginCmd.AddCommand(plugin_getCmd)

	carapace.Gen(plugin_getCmd).FlagCompletion(carapace.ActionMap{
		"arch": action.ActionImageArchitectures(),
		"os":   carapace.ActionValues("darwin", "linux", "windows"),
		"path": carapace.ActionDirectories(),
	})
}
