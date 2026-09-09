package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_config_setMirrorCmd = &cobra.Command{
	Use:   "set-mirror",
	Short: "Set a mirror for a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_config_setMirrorCmd).Standalone()
	package_config_setMirrorCmd.Flags().SetInterspersed(false)

	package_config_setMirrorCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_config_setMirrorCmd.Flags().String("mirror", "", "The mirror url or identity")
	package_config_setMirrorCmd.Flags().String("original", "", "The original url or identity")
	package_config_setMirrorCmd.Flags().Bool("version", false, "Show the version")

	package_configCmd.AddCommand(package_config_setMirrorCmd)
}
