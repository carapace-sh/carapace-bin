package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var installMultiPackageCmd = &cobra.Command{
	Use:   "install-multi-package",
	Short: "push one or more packages to the device and install them atomically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installMultiPackageCmd).Standalone()

	installMultiPackageCmd.Flags().BoolS("R", "R", false, "disallow replacing an existing application")
	installMultiPackageCmd.Flags().String("abi", "", "override platform's default ABI")
	installMultiPackageCmd.Flags().Bool("bypass-low-target-sdk-block", false, "bypass the minimum target SDK install block")
	installMultiPackageCmd.Flags().BoolS("d", "d", false, "allow version code downgrade (debuggable packages only)")
	installMultiPackageCmd.Flags().Bool("dont-kill", false, "do not kill the app if it is running")
	installMultiPackageCmd.Flags().Bool("force-queryable", false, "make the package visible to all other apps")
	installMultiPackageCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installMultiPackageCmd.Flags().StringS("i", "i", "", "set the installer package name")
	installMultiPackageCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installMultiPackageCmd.Flags().StringS("p", "p", "", "partial application install, inheriting the given installed package")
	installMultiPackageCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installMultiPackageCmd.Flags().BoolS("t", "t", false, "allow test packages")
	installMultiPackageCmd.Flags().String("user", "", "install for the given user")
	rootCmd.AddCommand(installMultiPackageCmd)

	carapace.Gen(installMultiPackageCmd).FlagCompletion(carapace.ActionMap{
		"abi":  carapace.ActionValues("arm64-v8a", "armeabi-v7a", "x86", "x86_64"),
		"i":    adb.ActionPackages(),
		"p":    adb.ActionPackages(),
		"user": actionInstallUsers(),
	})

	carapace.Gen(installMultiPackageCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".apk", ".apex"),
	)
}
