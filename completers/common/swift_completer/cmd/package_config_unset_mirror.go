package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_config_unsetMirrorCmd = &cobra.Command{
	Use:   "unset-mirror",
	Short: "Remove an existing mirror",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_config_unsetMirrorCmd).Standalone()
	package_config_unsetMirrorCmd.Flags().SetInterspersed(false)

	package_config_unsetMirrorCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_config_unsetMirrorCmd.Flags().String("mirror", "", "The mirror url or identity")
	package_config_unsetMirrorCmd.Flags().String("original", "", "The original url or identity")
	package_config_unsetMirrorCmd.Flags().Bool("version", false, "Show the version")

	package_configCmd.AddCommand(package_config_unsetMirrorCmd)
}
