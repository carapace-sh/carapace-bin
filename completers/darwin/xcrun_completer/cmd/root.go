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

	rootCmd.Flags().BoolP("find", "f", false, "Display the full path to the given tool")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("no-cache", "n", false, "Do not use the lookup cache")
	rootCmd.Flags().StringP("sdk", "sdk", "", "Use the given SDK")
	rootCmd.Flags().BoolP("show-sdk-build-version", "", false, "Show the SDK build version")
	rootCmd.Flags().BoolP("show-sdk-path", "", false, "Show the SDK path")
	rootCmd.Flags().BoolP("show-sdk-platform", "", false, "Show the SDK platform")
	rootCmd.Flags().BoolP("show-sdk-version", "", false, "Show the SDK version")
	rootCmd.Flags().BoolP("show-toolchain-version", "", false, "Show the toolchain version")
	rootCmd.Flags().StringP("toolchain", "toolchain", "", "Use the given toolchain")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	)
}
