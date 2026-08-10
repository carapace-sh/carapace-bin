package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var package_experimentalInstallCmd = &cobra.Command{
	Use:   "experimental-install",
	Short: "Install executable products of the current package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_experimentalInstallCmd).Standalone()
	package_experimentalInstallCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(package_experimentalInstallCmd)

	package_experimentalInstallCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_experimentalInstallCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_experimentalInstallCmd)
}
