package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "choco",
	Short: "Chocolatey package manager for Windows",
	Long:  "https://docs.chocolatey.org/en-us/choco/",
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
			"upgrade", "upgrade a package",
			"uninstall", "uninstall a package",
			"list", "list packages",
			"search", "search for packages",
			"info", "display package information",
			"outdated", "list outdated packages",
			"pin", "pin a package to prevent upgrades",
			"source", "manage package sources",
			"feature", "manage features",
			"config", "manage configuration",
			"new", "create a new package",
			"pack", "pack a nuspec into a nupkg",
			"push", "push a compiled nupkg to a source",
			"apikey", "set API key for a source",
		),
	)
}
