package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_macosFrameworkCmd = &cobra.Command{
	Use:   "macos-framework",
	Short: "Build macOS frameworks for use in existing Xcode projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_macosFrameworkCmd).Standalone()

	build_macosFrameworkCmd.Flags().Bool("debug", false, "Produce a framework for the debug build configuration.")
	build_macosFrameworkCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_macosFrameworkCmd.Flags().Bool("no-debug", false, "Do not produce a framework for the debug build configuration.")
	build_macosFrameworkCmd.Flags().Bool("no-profile", false, "Do not produce a framework for the profile build configuration.")
	build_macosFrameworkCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_macosFrameworkCmd.Flags().Bool("no-release", false, "Do not produce a framework for the release build configuration.")
	build_macosFrameworkCmd.Flags().StringP("output", "o", "", "Location to write the frameworks.")
	build_macosFrameworkCmd.Flags().Bool("profile", false, "Produce a framework for the profile build configuration.")
	build_macosFrameworkCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_macosFrameworkCmd.Flags().Bool("release", false, "Produce a framework for the release build configuration.")
	buildCmd.AddCommand(build_macosFrameworkCmd)

	carapace.Gen(build_macosFrameworkCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
