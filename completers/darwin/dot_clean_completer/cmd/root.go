package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dot_clean",
	Short: "merge ._* files with corresponding native files",
	Long:  "https://keith.github.io/xcode-manpages/dot_clean.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "Flat merge")
	rootCmd.Flags().String("keep", "", "Keep policy: mostrecent, dotbar, native")
	rootCmd.Flags().BoolS("m", "m", false, "Always delete dot underbar files")
	rootCmd.Flags().BoolS("n", "n", false, "Delete dot underbar file if there is no matching native file")
	rootCmd.Flags().BoolS("s", "s", false, "Follow symbolic links")
	rootCmd.Flags().BoolS("v", "v", false, "Print verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"keep": carapace.ActionValues("mostrecent", "dotbar", "native"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
