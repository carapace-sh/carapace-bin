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

	rootCmd.Flags().BoolP("all", "a", false, "Install all available updates")
	rootCmd.Flags().BoolP("available", "A", false, "List available updates only")
	rootCmd.Flags().BoolP("clear-catalog", "c", false, "Clear the local catalog")
	rootCmd.Flags().BoolP("download", "d", false, "Download updates without installing")
	rootCmd.Flags().BoolP("dump-state", "s", false, "Dump state to log")
	rootCmd.Flags().StringP("ignore", "i", "", "Ignore an update")
	rootCmd.Flags().StringP("install", "i", "", "Install an update")
	rootCmd.Flags().BoolP("list", "l", false, "List all available updates")
	rootCmd.Flags().BoolP("list-full", "L", false, "List all available updates with full details")
	rootCmd.Flags().BoolP("no-scan", "n", false, "Do not scan for updates")
	rootCmd.Flags().StringP("reset-ignored", "r", "", "Reset ignored updates")
	rootCmd.Flags().BoolP("restart", "r", false, "Restart after installation if required")
	rootCmd.Flags().BoolP("schedule", "S", false, "Set or check the scheduled update setting")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"--list", "List all available updates",
			"--install", "Install an update",
			"--download", "Download updates without installing",
			"--schedule", "Set or check the scheduled update setting",
			"--clear-catalog", "Clear the local catalog",
			"--ignore", "Ignore an update",
			"--reset-ignored", "Reset ignored updates",
		),
	)
}
