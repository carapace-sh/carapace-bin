package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_swiftPackageCmd = &cobra.Command{
	Use:   "swift-package",
	Short: "Build a Swift Package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_swiftPackageCmd).Standalone()

	build_swiftPackageCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_swiftPackageCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_swiftPackageCmd.Flags().Bool("no-debug", false, "Do not build a debug version of your app.")
	build_swiftPackageCmd.Flags().Bool("no-profile", false, "Do not build a version of your app specialized for performance profiling.")
	build_swiftPackageCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_swiftPackageCmd.Flags().Bool("no-release", false, "Do not build a release version of your app.")
	build_swiftPackageCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_swiftPackageCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_swiftPackageCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	buildCmd.AddCommand(build_swiftPackageCmd)
}
