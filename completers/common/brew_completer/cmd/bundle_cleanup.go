package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Uninstall all dependencies not present in the `Brewfile`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_cleanupCmd).Standalone()

	bundle_cleanupCmd.Flags().Bool("all", false, "Clean up all supported dependencies.")
	bundle_cleanupCmd.Flags().Bool("cargo", false, "Clean up Cargo packages.")
	bundle_cleanupCmd.Flags().Bool("cask", false, "Clean up Homebrew cask dependencies.")
	bundle_cleanupCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_cleanupCmd.Flags().Bool("flatpak", false, "Clean up Flatpak packages. Note: Linux only.")
	bundle_cleanupCmd.Flags().BoolP("force", "f", false, "Actually perform cleanup operations and reset Homebrew's global trust store to the `Brewfile` values.")
	bundle_cleanupCmd.Flags().Bool("formula", false, "Clean up Homebrew formula dependencies.")
	bundle_cleanupCmd.Flags().Bool("go", false, "Clean up Go packages.")
	bundle_cleanupCmd.Flags().Bool("help", false, "Show this message.")
	bundle_cleanupCmd.Flags().Bool("install", false, "Run `install` before cleaning up dependencies.")
	bundle_cleanupCmd.Flags().Bool("krew", false, "Clean up Krew plugins.")
	bundle_cleanupCmd.Flags().Bool("mas", false, "Clean up Mac App Store dependencies.")
	bundle_cleanupCmd.Flags().Bool("no-cargo", false, "`cleanup` without Cargo packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_CARGO` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cask", false, "Clean up without Homebrew cask dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_CASK` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-brew", false, "Clean up without Homebrew formula dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_BREW` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-cargo", false, "`cleanup` without Cargo packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_CARGO` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-cask", false, "Clean up without Homebrew cask dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_CASK` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-flatpak", false, "`cleanup` without Flatpak packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_FLATPAK` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-go", false, "`cleanup` without Go packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_GO` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-krew", false, "`cleanup` without Krew plugins. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_KREW` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-mas", false, "`cleanup` without Mac App Store dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_MAS` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-npm", false, "`cleanup` without npm packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_NPM` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-tap", false, "Clean up without Homebrew tap dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_TAP` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-uv", false, "`cleanup` without uv tools. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_UV` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-vscode", false, "`cleanup` without VSCode (and forks/variants) extensions. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_VSCODE` is set.")
	bundle_cleanupCmd.Flags().Bool("no-cleanup-winget", false, "`cleanup` without WinGet packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_WINGET` is set.")
	bundle_cleanupCmd.Flags().Bool("no-flatpak", false, "`cleanup` without Flatpak packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_FLATPAK` is set.")
	bundle_cleanupCmd.Flags().Bool("no-formula", false, "Clean up without Homebrew formula dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_BREW` is set.")
	bundle_cleanupCmd.Flags().Bool("no-go", false, "`cleanup` without Go packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_GO` is set.")
	bundle_cleanupCmd.Flags().Bool("no-krew", false, "`cleanup` without Krew plugins. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_KREW` is set.")
	bundle_cleanupCmd.Flags().Bool("no-mas", false, "`cleanup` without Mac App Store dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_MAS` is set.")
	bundle_cleanupCmd.Flags().Bool("no-npm", false, "`cleanup` without npm packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_NPM` is set.")
	bundle_cleanupCmd.Flags().Bool("no-tap", false, "Clean up without Homebrew tap dependencies. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_TAP` is set.")
	bundle_cleanupCmd.Flags().Bool("no-uv", false, "`cleanup` without uv tools. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_UV` is set.")
	bundle_cleanupCmd.Flags().Bool("no-vscode", false, "`cleanup` without VSCode (and forks/variants) extensions. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_VSCODE` is set.")
	bundle_cleanupCmd.Flags().Bool("no-winget", false, "`cleanup` without WinGet packages. Enabled by default if `$HOMEBREW_BUNDLE_CLEANUP_NO_WINGET` is set.")
	bundle_cleanupCmd.Flags().Bool("npm", false, "Clean up npm packages.")
	bundle_cleanupCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_cleanupCmd.Flags().Bool("tap", false, "Clean up Homebrew tap dependencies.")
	bundle_cleanupCmd.Flags().Bool("uv", false, "Clean up uv tools.")
	bundle_cleanupCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundle_cleanupCmd.Flags().Bool("vscode", false, "Clean up VSCode (and forks/variants) extensions.")
	bundle_cleanupCmd.Flags().Bool("winget", false, "Clean up WinGet packages. Note: WSL only.")
	bundle_cleanupCmd.Flags().Bool("zap", false, "Clean up casks using the `zap` command instead of `uninstall`.")
	bundleCmd.AddCommand(bundle_cleanupCmd)
}
