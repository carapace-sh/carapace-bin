package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bundle_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove entries that match `name` from your `Brewfile`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_removeCmd).Standalone()

	bundle_removeCmd.Flags().Bool("cargo", false, "Remove entries for Cargo packages.")
	bundle_removeCmd.Flags().Bool("cask", false, "Remove Homebrew cask entries.")
	bundle_removeCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_removeCmd.Flags().Bool("flatpak", false, "Remove entries for Flatpak packages. Note: Linux only.")
	bundle_removeCmd.Flags().Bool("formula", false, "Remove Homebrew formula entries, including matches against formula aliases and old names.")
	bundle_removeCmd.Flags().Bool("go", false, "Remove entries for Go packages.")
	bundle_removeCmd.Flags().Bool("help", false, "Show this message.")
	bundle_removeCmd.Flags().Bool("install", false, "Run `install` before removing entries.")
	bundle_removeCmd.Flags().Bool("krew", false, "Remove entries for Krew plugins.")
	bundle_removeCmd.Flags().Bool("mas", false, "Remove entries for Mac App Store dependencies.")
	bundle_removeCmd.Flags().Bool("npm", false, "Remove entries for npm packages.")
	bundle_removeCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_removeCmd.Flags().Bool("tap", false, "Remove Homebrew tap entries.")
	bundle_removeCmd.Flags().Bool("uv", false, "Remove entries for uv tools.")
	bundle_removeCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundle_removeCmd.Flags().Bool("vscode", false, "Remove entries for VSCode (and forks/variants) extensions.")
	bundle_removeCmd.Flags().Bool("winget", false, "Remove entries for WinGet packages. Note: WSL only.")
	bundleCmd.AddCommand(bundle_removeCmd)

	carapace.Gen(bundle_removeCmd).PositionalAnyCompletion(
		action.ActionBundleEntries(bundle_removeCmd).FilterArgs(),
	)
}
