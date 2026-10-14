package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all dependencies present in the `Brewfile`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_listCmd).Standalone()

	bundle_listCmd.Flags().Bool("all", false, "List all dependencies.")
	bundle_listCmd.Flags().Bool("cargo", false, "List Cargo packages.")
	bundle_listCmd.Flags().Bool("cask", false, "List Homebrew cask dependencies.")
	bundle_listCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_listCmd.Flags().Bool("flatpak", false, "List Flatpak packages. Note: Linux only.")
	bundle_listCmd.Flags().Bool("formula", false, "List Homebrew formula dependencies.")
	bundle_listCmd.Flags().Bool("go", false, "List Go packages.")
	bundle_listCmd.Flags().Bool("help", false, "Show this message.")
	bundle_listCmd.Flags().Bool("install", false, "Run `install` before listing dependencies.")
	bundle_listCmd.Flags().Bool("krew", false, "List Krew plugins.")
	bundle_listCmd.Flags().Bool("mas", false, "List Mac App Store dependencies.")
	bundle_listCmd.Flags().Bool("npm", false, "List npm packages.")
	bundle_listCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_listCmd.Flags().Bool("tap", false, "List Homebrew tap dependencies.")
	bundle_listCmd.Flags().Bool("uv", false, "List uv tools.")
	bundle_listCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundle_listCmd.Flags().Bool("vscode", false, "List VSCode (and forks/variants) extensions.")
	bundle_listCmd.Flags().Bool("winget", false, "List WinGet packages. Note: WSL only.")
	bundleCmd.AddCommand(bundle_listCmd)
}
