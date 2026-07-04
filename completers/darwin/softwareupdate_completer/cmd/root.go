package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "softwareupdate",
	Short: "system software update tool",
	Long:  "https://keith.github.io/xcode-manpages/softwareupdate.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Install all available updates")
	rootCmd.Flags().Bool("background", false, "Background mode")
	rootCmd.Flags().Bool("download", false, "Download updates without installing")
	rootCmd.Flags().Bool("fetch-full-installer", false, "Fetch the full macOS installer")
	rootCmd.Flags().Bool("help", false, "Display usage information")
	rootCmd.Flags().Bool("install", false, "Install updates")
	rootCmd.Flags().Bool("install-rosetta", false, "Install Rosetta")
	rootCmd.Flags().Bool("list", false, "List all available updates")
	rootCmd.Flags().Bool("list-full-installers", false, "List all available full installers")
	rootCmd.Flags().Bool("os-only", false, "Only install OS updates")
	rootCmd.Flags().Bool("recommended", false, "Install recommended updates only")
	rootCmd.Flags().Bool("restart", false, "Restart after installation if required")
	rootCmd.Flags().Bool("safari-only", false, "Only install Safari updates")
	rootCmd.Flags().Bool("schedule", false, "Set or check the scheduled update setting")
	rootCmd.Flags().Bool("stdinpass", false, "Read password from stdin")
	rootCmd.Flags().Bool("user", false, "Interactive mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"--list", "List all available updates",
			"--install", "Install updates",
			"--download", "Download updates without installing",
			"--schedule", "Set or check the scheduled update setting",
			"--list-full-installers", "List all available full installers",
			"--fetch-full-installer", "Fetch the full macOS installer",
			"--install-rosetta", "Install Rosetta",
			"--background", "Background mode",
			"--help", "Display usage information",
		),
	)
}
