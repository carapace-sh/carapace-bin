package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var AddProvisionedAppxPackageCmd = &cobra.Command{
	Use:   "Add-ProvisionedAppxPackage",
	Short: "add app packages to an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(AddProvisionedAppxPackageCmd).Standalone()
	rootCmd.AddCommand(AddProvisionedAppxPackageCmd)
}
