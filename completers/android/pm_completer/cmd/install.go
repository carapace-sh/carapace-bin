package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("abi", "", "override the default `ABI` of the platform")
	installCmd.Flags().BoolP("allow-downgrade", "d", false, "allow version code downgrade (debuggable packages only)")
	installCmd.Flags().Bool("apex", false, "install an .apex file, not an .apk")
	installCmd.Flags().Bool("dont-kill", false, "installing a new feature split, don't kill running app")
	installCmd.Flags().Bool("enable-rollback", false, "enable rollbacks for the upgrade")
	installCmd.Flags().String("force-uuid", "", "force install on to disk volume with given `UUID`")
	installCmd.Flags().Bool("full", false, "cause the app to be installed as a non-ephemeral full app")
	installCmd.Flags().BoolP("grant-permissions", "g", false, "grant all runtime permissions")
	installCmd.Flags().String("install-location", "", "force the install `LOCATION` (0=auto, 1=internal, 2=prefer external)")
	installCmd.Flags().String("install-reason", "", "indicates `REASON` why the app is being installed (0-4)")
	installCmd.Flags().StringP("installer", "i", "", "package name of installer owning the app")
	installCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installCmd.Flags().BoolP("internal", "f", false, "install application on internal flash")
	installCmd.Flags().BoolP("no-replace", "R", false, "disallow replacement of existing application")
	installCmd.Flags().Bool("non-staged", false, "explicitly set this installation to be non-staged")
	installCmd.Flags().String("originating-uri", "", "set `URI` where app was downloaded from")
	installCmd.Flags().StringP("package-size", "S", "", "size in `BYTES` of package, required for stdin")
	installCmd.Flags().BoolP("partial", "p", false, "partial application install (new split on top of existing pkg)")
	installCmd.Flags().String("pkg", "", "specify expected package name of app being installed")
	installCmd.Flags().Bool("preload", false, "install as a preloaded app")
	installCmd.Flags().String("referrer", "", "set `URI` that instigated the install of the app")
	installCmd.Flags().Bool("restrict-permissions", false, "don't whitelist restricted permissions at install")
	installCmd.Flags().BoolP("test", "t", false, "allow test packages")
	installCmd.Flags().String("user", "", "install under the given `USER_ID`")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"install-location": carapace.ActionValues("0", "1", "2"),
		"install-reason":   carapace.ActionValues("0", "1", "2", "3", "4"),
		"installer":        carapace.ActionFiles(),
		"originating-uri":  carapace.ActionValues(),
		"pkg":              carapace.ActionValues(),
		"referrer":         carapace.ActionValues(),
		"user": carapace.Batch(
			carapace.ActionValues("all", "current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
