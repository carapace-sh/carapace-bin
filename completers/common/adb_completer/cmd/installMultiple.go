package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var installMultipleCmd = &cobra.Command{
	Use:   "install-multiple",
	Short: "push multiple APKs to the device for a single package and install them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installMultipleCmd).Standalone()

	installMultipleCmd.Flags().BoolS("R", "R", false, "disallow replacing an existing application")
	installMultipleCmd.Flags().String("abi", "", "override platform's default ABI")
	installMultipleCmd.Flags().Bool("bypass-low-target-sdk-block", false, "bypass the minimum target SDK install block")
	installMultipleCmd.Flags().BoolS("d", "d", false, "allow version code downgrade (debuggable packages only)")
	installMultipleCmd.Flags().Bool("dont-kill", false, "do not kill the app if it is running")
	installMultipleCmd.Flags().Bool("force-queryable", false, "make the package visible to all other apps")
	installMultipleCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installMultipleCmd.Flags().StringS("i", "i", "", "set the installer package name")
	installMultipleCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installMultipleCmd.Flags().Bool("no-streaming", false, "always push APK to device and invoke Package Manager as separate steps")
	installMultipleCmd.Flags().StringS("p", "p", "", "partial application install, inheriting the given installed package")
	installMultipleCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installMultipleCmd.Flags().Bool("streaming", false, "force streaming APK directly into Package Manager")
	installMultipleCmd.Flags().BoolS("t", "t", false, "allow test packages")
	installMultipleCmd.Flags().String("user", "", "install for the given user")
	rootCmd.AddCommand(installMultipleCmd)

	carapace.Gen(installMultipleCmd).FlagCompletion(carapace.ActionMap{
		"abi":  carapace.ActionValues("arm64-v8a", "armeabi-v7a", "x86", "x86_64"),
		"i":    adb.ActionPackages(),
		"p":    adb.ActionPackages(),
		"user": actionInstallUsers(),
	})

	carapace.Gen(installMultipleCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".apk"),
	)
}
