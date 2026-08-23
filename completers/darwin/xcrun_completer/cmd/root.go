package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xcrun",
	Short: "run or locate development commands",
	Long:  "https://keith.github.io/xcode-manpages/xcrun.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("find", "f", false, "Only find and print the tool path")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("kill-cache", "k", false, "Invalidate all existing cache entries")
	rootCmd.Flags().BoolP("log", "l", false, "Show commands to be executed (with --run)")
	rootCmd.Flags().BoolP("no-cache", "n", false, "Do not use the lookup cache")
	rootCmd.Flags().BoolP("run", "r", false, "Find and execute the tool (default)")
	rootCmd.Flags().String("sdk", "", "Use the given SDK")
	rootCmd.Flags().Bool("show-sdk-build-version", false, "Show selected SDK build version")
	rootCmd.Flags().Bool("show-sdk-path", false, "Show selected SDK install path")
	rootCmd.Flags().Bool("show-sdk-platform-path", false, "Show selected SDK platform path")
	rootCmd.Flags().Bool("show-sdk-platform-version", false, "Show selected SDK platform version")
	rootCmd.Flags().Bool("show-sdk-version", false, "Show selected SDK version")
	rootCmd.Flags().Bool("show-toolchain-path", false, "Show selected toolchain path")
	rootCmd.Flags().String("toolchain", "", "Find the tool for the given toolchain")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
	rootCmd.Flags().Bool("version", false, "Show the xcrun version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"sdk":       carapace.ActionValues("macosx", "iphoneos", "iphonesimulator", "appletvos", "appletvsimulator", "watchos", "watchsimulator", "xros", "xrsimulator"),
		"toolchain": carapace.ActionValues("default", "com.apple.dt.toolchain.XcodeDefault", "swift", "org.swift.5220190702a"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	)
}
