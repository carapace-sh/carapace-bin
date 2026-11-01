package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Write all installed casks/formulae/images/taps into a `Brewfile`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_dumpCmd).Standalone()

	bundle_dumpCmd.Flags().Bool("cargo", false, "Dump Cargo packages.")
	bundle_dumpCmd.Flags().Bool("cask", false, "Dump Homebrew cask dependencies.")
	bundle_dumpCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_dumpCmd.Flags().Bool("flatpak", false, "Dump Flatpak packages. Note: Linux only.")
	bundle_dumpCmd.Flags().BoolP("force", "f", false, "Overwrite an existing `Brewfile`.")
	bundle_dumpCmd.Flags().Bool("formula", false, "Dump Homebrew formula dependencies.")
	bundle_dumpCmd.Flags().Bool("go", false, "Dump Go packages.")
	bundle_dumpCmd.Flags().Bool("help", false, "Show this message.")
	bundle_dumpCmd.Flags().Bool("install", false, "Run `install` before dumping dependencies.")
	bundle_dumpCmd.Flags().Bool("krew", false, "Dump Krew plugins.")
	bundle_dumpCmd.Flags().Bool("mas", false, "Dump Mac App Store dependencies.")
	bundle_dumpCmd.Flags().Bool("no-cargo", false, "`dump` without Cargo packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_CARGO` is set.")
	bundle_dumpCmd.Flags().Bool("no-cask", false, "Dump without Homebrew cask dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_CASK` is set.")
	bundle_dumpCmd.Flags().Bool("no-describe", false, "Do not add description comments above each line. Description comments are the default. Enabled by default if `$HOMEBREW_BUNDLE_NO_DESCRIBE` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-brew", false, "Dump without Homebrew formula dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_BREW` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-cargo", false, "`dump` without Cargo packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_CARGO` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-cask", false, "Dump without Homebrew cask dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_CASK` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-flatpak", false, "`dump` without Flatpak packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_FLATPAK` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-go", false, "`dump` without Go packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_GO` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-krew", false, "`dump` without Krew plugins. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_KREW` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-mas", false, "`dump` without Mac App Store dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_MAS` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-npm", false, "`dump` without npm packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_NPM` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-tap", false, "Dump without Homebrew tap dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_TAP` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-uv", false, "`dump` without uv tools. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_UV` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-vscode", false, "`dump` without VSCode (and forks/variants) extensions. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_VSCODE` is set.")
	bundle_dumpCmd.Flags().Bool("no-dump-winget", false, "`dump` without WinGet packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_WINGET` is set.")
	bundle_dumpCmd.Flags().Bool("no-flatpak", false, "`dump` without Flatpak packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_FLATPAK` is set.")
	bundle_dumpCmd.Flags().Bool("no-formula", false, "Dump without Homebrew formula dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_BREW` is set.")
	bundle_dumpCmd.Flags().Bool("no-go", false, "`dump` without Go packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_GO` is set.")
	bundle_dumpCmd.Flags().Bool("no-krew", false, "`dump` without Krew plugins. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_KREW` is set.")
	bundle_dumpCmd.Flags().Bool("no-mas", false, "`dump` without Mac App Store dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_MAS` is set.")
	bundle_dumpCmd.Flags().Bool("no-npm", false, "`dump` without npm packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_NPM` is set.")
	bundle_dumpCmd.Flags().Bool("no-restart", false, "Do not add `restart_service` to formula lines.")
	bundle_dumpCmd.Flags().Bool("no-tap", false, "Dump without Homebrew tap dependencies. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_TAP` is set.")
	bundle_dumpCmd.Flags().Bool("no-uv", false, "`dump` without uv tools. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_UV` is set.")
	bundle_dumpCmd.Flags().Bool("no-vscode", false, "`dump` without VSCode (and forks/variants) extensions. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_VSCODE` is set.")
	bundle_dumpCmd.Flags().Bool("no-winget", false, "`dump` without WinGet packages. Enabled by default if `$HOMEBREW_BUNDLE_DUMP_NO_WINGET` is set.")
	bundle_dumpCmd.Flags().Bool("npm", false, "Dump npm packages.")
	bundle_dumpCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_dumpCmd.Flags().Bool("tap", false, "Dump Homebrew tap dependencies.")
	bundle_dumpCmd.Flags().Bool("uv", false, "Dump uv tools.")
	bundle_dumpCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundle_dumpCmd.Flags().Bool("vscode", false, "Dump VSCode (and forks/variants) extensions.")
	bundle_dumpCmd.Flags().Bool("winget", false, "Dump WinGet packages. Note: WSL only.")
	bundleCmd.AddCommand(bundle_dumpCmd)
}
