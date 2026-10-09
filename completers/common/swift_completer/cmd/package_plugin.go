package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Invoke a command plugin or perform other actions on command plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_pluginCmd).Standalone()
	package_pluginCmd.Flags().SetInterspersed(false)

	package_pluginCmd.Flags().String("allow-network-connections", "", "Allow the plugin to make network connections")
	package_pluginCmd.Flags().StringArray("allow-writing-to-directory", nil, "Allow the plugin to write to an additional directory")
	package_pluginCmd.Flags().Bool("allow-writing-to-package-directory", false, "Allow the plugin to write to the package directory")
	package_pluginCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_pluginCmd.Flags().Bool("list", false, "List the available command plugins")
	package_pluginCmd.Flags().String("package", "", "Limit available plugins to a single package with the given identity")
	package_pluginCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_pluginCmd)

	carapace.Gen(package_pluginCmd).FlagCompletion(carapace.ActionMap{
		"allow-network-connections":  carapace.ActionValues("none", "local", "all", "docker", "unixDomainSocket"),
		"allow-writing-to-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(package_pluginCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
