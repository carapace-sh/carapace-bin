package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var package_experimentalUninstallCmd = &cobra.Command{
	Use:   "experimental-uninstall",
	Short: "Uninstall executable products previously installed by experimental-install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_experimentalUninstallCmd).Standalone()
	package_experimentalUninstallCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(package_experimentalUninstallCmd)

	package_experimentalUninstallCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_experimentalUninstallCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_experimentalUninstallCmd)
}
