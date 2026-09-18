package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_config_getMirrorCmd = &cobra.Command{
	Use:   "get-mirror",
	Short: "Print mirror configuration for the given package dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_config_getMirrorCmd).Standalone()
	package_config_getMirrorCmd.Flags().SetInterspersed(false)

	package_config_getMirrorCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_config_getMirrorCmd.Flags().String("original", "", "The original url or identity")
	package_config_getMirrorCmd.Flags().Bool("version", false, "Show the version")

	package_configCmd.AddCommand(package_config_getMirrorCmd)
}
