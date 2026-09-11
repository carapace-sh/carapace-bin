package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bundle_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add entries to your `Brewfile`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_addCmd).Standalone()

	bundle_addCmd.Flags().Bool("cargo", false, "Add entries for Cargo packages.")
	bundle_addCmd.Flags().Bool("cask", false, "Add Homebrew cask entries.")
	bundle_addCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_addCmd.Flags().Bool("flatpak", false, "Add entries for Flatpak packages. Note: Linux only.")
	bundle_addCmd.Flags().Bool("formula", false, "Add Homebrew formula entries.")
	bundle_addCmd.Flags().Bool("go", false, "Add entries for Go packages.")
	bundle_addCmd.Flags().Bool("help", false, "Show this message.")
	bundle_addCmd.Flags().Bool("install", false, "Run `install` before adding entries.")
	bundle_addCmd.Flags().Bool("krew", false, "Add entries for Krew plugins.")
	bundle_addCmd.Flags().Bool("no-describe", false, "Do not add description comments above each line. Description comments are the default. Enabled by default if `$HOMEBREW_BUNDLE_NO_DESCRIBE` is set.")
	bundle_addCmd.Flags().Bool("npm", false, "Add entries for npm packages.")
	bundle_addCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_addCmd.Flags().Bool("tap", false, "Add Homebrew tap entries.")
	bundle_addCmd.Flags().Bool("uv", false, "Add entries for uv tools.")
	bundle_addCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundle_addCmd.Flags().Bool("vscode", false, "Add entries for VSCode (and forks/variants) extensions.")
	bundleCmd.AddCommand(bundle_addCmd)

	carapace.Gen(bundle_addCmd).PositionalAnyCompletion(
		action.ActionBundlePackages(bundle_addCmd).FilterArgs(),
	)
}
