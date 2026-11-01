package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var precacheCmd = &cobra.Command{
	Use:   "precache",
	Short: "Populate the Flutter tool's cache of binary artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(precacheCmd).Standalone()

	precacheCmd.Flags().Bool("all", false, "Precache artifacts for all host platforms.")
	precacheCmd.Flags().BoolP("all-platforms", "a", false, "Precache artifacts for all host platforms.")
	precacheCmd.Flags().Bool("android", false, "Precache artifacts for Android development.")
	precacheCmd.Flags().Bool("android_gen_snapshot", false, "Precache gen_snapshot for Android development.")
	precacheCmd.Flags().Bool("android_internal_build", false, "Precache dependencies for internal Android development.")
	precacheCmd.Flags().Bool("android_maven", false, "Precache Gradle dependencies for Android development.")
	precacheCmd.Flags().Bool("flutter_runner", false, "Precache the flutter runner artifacts.")
	precacheCmd.Flags().BoolP("force", "f", false, "Force re-downloading of artifacts.")
	precacheCmd.Flags().Bool("fuchsia", false, "Precache artifacts for Fuchsia development.")
	precacheCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	precacheCmd.Flags().String("host-arch", "", "Override the architecture of host artifacts to precache.")
	precacheCmd.Flags().Bool("informative", false, "Precache engine_stamp.json.")
	precacheCmd.Flags().Bool("ios", false, "Precache artifacts for iOS development.")
	precacheCmd.Flags().Bool("linux", false, "Precache artifacts for Linux desktop development.")
	precacheCmd.Flags().Bool("macos", false, "Precache artifacts for macOS desktop development.")
	precacheCmd.Flags().Bool("no-android", false, "Precache artifacts for Android development.")
	precacheCmd.Flags().Bool("no-android_gen_snapshot", false, "Precache gen_snapshot for Android development.")
	precacheCmd.Flags().Bool("no-android_internal_build", false, "Precache dependencies for internal Android development.")
	precacheCmd.Flags().Bool("no-android_maven", false, "Precache Gradle dependencies for Android development.")
	precacheCmd.Flags().Bool("no-flutter_runner", false, "Precache the flutter runner artifacts.")
	precacheCmd.Flags().Bool("no-fuchsia", false, "Precache artifacts for Fuchsia development.")
	precacheCmd.Flags().Bool("no-ios", false, "Precache artifacts for iOS development.")
	precacheCmd.Flags().Bool("no-linux", false, "Precache artifacts for Linux desktop development.")
	precacheCmd.Flags().Bool("no-macos", false, "Precache artifacts for macOS desktop development.")
	precacheCmd.Flags().Bool("no-universal", false, "Precache artifacts required for any development platform.")
	precacheCmd.Flags().Bool("no-use-unsigned-mac-binaries", false, "Precache the unsigned macOS binaries when available.")
	precacheCmd.Flags().Bool("no-web", false, "Precache artifacts for web development.")
	precacheCmd.Flags().Bool("no-windows", false, "Precache artifacts for Windows desktop development.")
	precacheCmd.Flags().Bool("universal", false, "Precache artifacts required for any development platform.")
	precacheCmd.Flags().Bool("use-unsigned-mac-binaries", false, "Precache the unsigned macOS binaries when available.")
	precacheCmd.Flags().Bool("web", false, "Precache artifacts for web development.")
	precacheCmd.Flags().Bool("windows", false, "Precache artifacts for Windows desktop development.")
	rootCmd.AddCommand(precacheCmd)
	carapace.Gen(precacheCmd).FlagCompletion(carapace.ActionMap{
		"host-arch": carapace.ActionValues("arm64", "x64"),
	})

}
