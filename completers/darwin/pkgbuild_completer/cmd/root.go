package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkgbuild",
	Short: "build a macOS Installer component package",
	Long:  "https://keith.github.io/xcode-manpages/pkgbuild.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("analyze", false, "Analyze mode")
	rootCmd.Flags().String("component", "", "Component path")
	rootCmd.Flags().String("component-plist", "", "Component plist path")
	rootCmd.Flags().String("root", "", "Root path")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"component":       carapace.ActionFiles(),
		"component-plist": carapace.ActionFiles(),
		"root":            carapace.ActionDirectories(),
	})
}
