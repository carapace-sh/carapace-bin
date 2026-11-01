package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete the build/ and .dart_tool/ directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	cleanCmd.Flags().Bool("include-example", false, "Also clean the example directory, if one exists. Useful when developing in a package project.")
	cleanCmd.Flags().Bool("include-xcode-workspace", false, "Whether to run \"xcodebuild clean\" on the Xcode workspace for iOS and macOS projects. This removes build products and intermediate files from Xcode's build cache and can be slow to complete.")
	cleanCmd.Flags().String("scheme", "", "When cleaning Xcode schemes, clean only the specified scheme.")
	cleanCmd.Flags().Bool("stop-gradle", false, "Force stop active Gradle daemons before or during clean. Useful on Windows when files in build/ are locked by background processes.")
	rootCmd.AddCommand(cleanCmd)
}
