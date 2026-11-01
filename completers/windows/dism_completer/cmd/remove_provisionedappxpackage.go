package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemoveProvisionedAppxPackageCmd = &cobra.Command{
	Use:   "Remove-ProvisionedAppxPackage",
	Short: "remove provisioning for app packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemoveProvisionedAppxPackageCmd).Standalone()
	rootCmd.AddCommand(RemoveProvisionedAppxPackageCmd)
}
