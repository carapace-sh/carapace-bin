package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addTargetPluginCmd = &cobra.Command{
	Use:   "add-target-plugin",
	Short: "Add a new target plugin to the manifest",
	Args:  cobra.ExactArgs(2),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addTargetPluginCmd).Standalone()
	package_addTargetPluginCmd.Flags().SetInterspersed(false)

	package_addTargetPluginCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addTargetPluginCmd.Flags().String("package", "", "The package in which the plugin resides")
	package_addTargetPluginCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addTargetPluginCmd)
}
