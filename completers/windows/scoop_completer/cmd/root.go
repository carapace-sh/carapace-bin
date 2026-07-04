package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scoop",
	Short: "command-line installer for Windows",
	Long:  "https://scoop.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"install", "install a package",
			"uninstall", "uninstall a package",
			"update", "update packages or scoop itself",
			"list", "list installed packages",
			"search", "search for packages",
			"info", "show package information",
			"cache", "manage download cache",
			"cleanup", "cleanup old versions",
			"config", "manage configuration",
			"home", "open package homepage",
			"which", "find which shim an executable resolves to",
			"shim", "manage shims",
			"reset", "reset a package",
			"bucket", "manage buckets",
		),
	)
}
