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

	rootCmd.Flags().Bool("analyze", false, "Analyze the target")
	rootCmd.Flags().String("arch", "", "Build for the specified architecture")
	rootCmd.Flags().Bool("archive", false, "Archive the target")
	rootCmd.Flags().Bool("build", false, "Build the target")
	rootCmd.Flags().Bool("clean", false, "Clean the build directory")
	rootCmd.Flags().String("configuration", "", "Use the specified build configuration")
	rootCmd.Flags().String("derivedDataPath", "", "Override the derived data path")
	rootCmd.Flags().String("destination", "", "Use the specified destination")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().Bool("list", false, "List targets and schemes")
	rootCmd.Flags().String("project", "", "Build the specified project")
	rootCmd.Flags().String("resultBundlePath", "", "Write result bundle to the specified path")
	rootCmd.Flags().String("scheme", "", "Build the specified scheme")
	rootCmd.Flags().String("sdk", "", "Use the specified SDK")
	rootCmd.Flags().Bool("showBuildSettings", false, "Show build settings")
	rootCmd.Flags().Bool("showsdks", false, "Show available SDKs")
	rootCmd.Flags().String("target", "", "Build the specified target")
	rootCmd.Flags().Bool("test", false, "Run the test action")
	rootCmd.Flags().String("workspace", "", "Build the specified workspace")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"arch":            carapace.ActionValues("arm64", "x86_64"),
		"configuration":   carapace.ActionValues("Debug", "Release"),
		"derivedDataPath": carapace.ActionDirectories(),
		"destination": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"generic/platform=iOS", "iOS",
						"generic/platform=macOS", "macOS",
						"generic/platform=tvOS", "tvOS",
						"generic/platform=watchOS", "watchOS",
						"platform=iOS Simulator", "iOS Simulator",
						"platform=macOS,variant=Mac Catalyst", "macOS Catalyst",
					)
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"project":          carapace.ActionFiles(".xcodeproj"),
		"resultBundlePath": carapace.ActionFiles(),
		"sdk":              carapace.ActionValues("macosx", "iphoneos", "iphonesimulator", "appletvos", "appletvsimulator", "watchos", "watchsimulator", "xros", "xrsimulator"),
		"workspace":        carapace.ActionFiles(".xcworkspace"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
