package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "productbuild",
	Short: "build a product archive for the macOS Installer",
	Long:  "https://keith.github.io/xcode-manpages/productbuild.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("component", "", "Component path")
	rootCmd.Flags().String("content", "", "Content path")
	rootCmd.Flags().String("distribution", "", "Distribution path")
	rootCmd.Flags().String("package", "", "Package path")
	rootCmd.Flags().String("package-path", "", "Package search path")
	rootCmd.Flags().String("product", "", "Requirements plist")
	rootCmd.Flags().String("root", "", "Root path")
	rootCmd.Flags().Bool("synthesize", false, "Synthesize distribution")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"component":    carapace.ActionFiles(),
		"content":      carapace.ActionFiles(),
		"distribution": carapace.ActionFiles(),
		"package":      carapace.ActionFiles(),
		"package-path": carapace.ActionDirectories(),
		"product":      carapace.ActionFiles(),
		"root":         carapace.ActionDirectories(),
	})
}
