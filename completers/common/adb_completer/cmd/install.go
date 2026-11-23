package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "push a single package to the device and install it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolS("R", "R", false, "disallow replacing an existing application")
	installCmd.Flags().String("abi", "", "override platform's default ABI")
	installCmd.Flags().Bool("bypass-low-target-sdk-block", false, "bypass the minimum target SDK install block")
	installCmd.Flags().BoolS("d", "d", false, "allow version code downgrade (debuggable packages only)")
	installCmd.Flags().Bool("date-check-agent", false, "update deployment agent when local version is newer and using fast deploy")
	installCmd.Flags().Bool("dont-kill", false, "do not kill the app if it is running")
	installCmd.Flags().Bool("fastdeploy", false, "use fast deploy")
	installCmd.Flags().Bool("force-agent", false, "force update of deployment agent when using fast deploy")
	installCmd.Flags().Bool("force-queryable", false, "make the package visible to all other apps")
	installCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installCmd.Flags().StringS("i", "i", "", "set the installer package name")
	installCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installCmd.Flags().Bool("no-fastdeploy", false, "prevent use of fast deploy")
	installCmd.Flags().Bool("no-streaming", false, "always push APK to device and invoke Package Manager as separate steps")
	installCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installCmd.Flags().Bool("streaming", false, "force streaming APK directly into Package Manager")
	installCmd.Flags().BoolS("t", "t", false, "allow test packages")
	installCmd.Flags().String("user", "", "install for the given user")
	installCmd.Flags().Bool("version-check-agent", false, "update deployment agent when local version has different version code and using fast deploy")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"abi":  carapace.ActionValues("arm64-v8a", "armeabi-v7a", "x86", "x86_64"),
		"i":    adb.ActionPackages(),
		"user": actionInstallUsers(),
	})

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionFiles(".apk"),
	)
}

// actionInstallUsers completes the user argument Package Manager accepts:
// a device user id, or `all` / `current`.
func actionInstallUsers() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"all", "all users",
			"current", "the current foreground user",
		),
		adb.ActionUsers(),
	).ToA()
}
