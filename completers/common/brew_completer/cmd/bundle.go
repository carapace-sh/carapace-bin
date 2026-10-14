package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:     "bundle",
	Short:   "Bundler for non-Ruby dependencies from Homebrew, Homebrew Cask, Mac App Store, Visual Studio Code and more",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()

	bundleCmd.PersistentFlags().String("file", "", "Read from or write to the `Brewfile` from this location. Use `--file=-` to pipe to stdin/stdout.")
	bundleCmd.PersistentFlags().BoolP("global", "g", false, "Read from or write to the `Brewfile` from `$HOMEBREW_BUNDLE_FILE_GLOBAL` (if set), `${XDG_CONFIG_HOME}/homebrew/Brewfile` (if `$XDG_CONFIG_HOME` is set), `~/.homebrew/Brewfile` or `~/.Brewfile` otherwise.")

	bundleCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundleCmd.Flags().Bool("help", false, "Show this message.")
	bundleCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundleCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(bundleCmd)

	carapace.Gen(bundleCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
