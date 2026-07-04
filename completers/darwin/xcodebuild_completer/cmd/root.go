package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xcodebuild",
	Short: "build Xcode projects and workspaces",
	Long:  "https://keith.github.io/xcode-manpages/xcodebuild.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("analyze", "analyze", false, "Analyze the target")
	rootCmd.Flags().StringP("arch", "arch", "", "Build for the specified architecture")
	rootCmd.Flags().BoolP("archive", "archive", false, "Archive the target")
	rootCmd.Flags().BoolP("build", "build", false, "Build the target")
	rootCmd.Flags().BoolP("clean", "clean", false, "Clean the build directory")
	rootCmd.Flags().StringP("configuration", "configuration", "", "Use the specified build configuration")
	rootCmd.Flags().StringP("derivedDataPath", "derivedDataPath", "", "Override the derived data path")
	rootCmd.Flags().StringP("destination", "destination", "", "Use the specified destination")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("list", "list", false, "List targets and schemes")
	rootCmd.Flags().StringP("project", "project", "", "Build the specified project")
	rootCmd.Flags().StringP("resultBundlePath", "resultBundlePath", "", "Write result bundle to the specified path")
	rootCmd.Flags().StringP("scheme", "scheme", "", "Build the specified scheme")
	rootCmd.Flags().StringP("sdk", "sdk", "", "Use the specified SDK")
	rootCmd.Flags().BoolP("showBuildSettings", "showBuildSettings", false, "Show build settings")
	rootCmd.Flags().BoolP("showsdks", "showsdks", false, "Show available SDKs")
	rootCmd.Flags().StringP("target", "target", "", "Build the specified target")
	rootCmd.Flags().BoolP("test", "test", false, "Run the test action")
	rootCmd.Flags().StringP("workspace", "workspace", "", "Build the specified workspace")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionValues("Debug", "Release"),
		"project":       carapace.ActionFiles(".xcodeproj"),
		"workspace":     carapace.ActionFiles(".xcworkspace"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
