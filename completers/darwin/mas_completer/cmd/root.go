package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mas",
	Short: "Mac App Store command line interface",
	Long:  "https://github.com/mas-cli/mas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"account", "Print account information",
			"home", "Open Mac App Store page for app",
			"info", "Display app information",
			"install", "Install an app",
			"list", "List installed apps",
			"lucky", "Install first search result",
			"open", "Open Mac App Store",
			"outdated", "List outdated apps",
			"purchase", "Purchase and install an app",
			"reset", "Reset Mac App Store",
			"search", "Search for apps",
			"signin", "Sign in to Mac App Store",
			"signout", "Sign out of Mac App Store",
			"uninstall", "Uninstall an app",
			"upgrade", "Upgrade outdated apps",
			"vendor", "Open vendor page for app",
			"version", "Print version",
		),
	)
}
