package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systemextensionsctl",
	Short: "Manage system extensions",
	Long:  "https://developer.apple.com/documentation/systemextensions",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"developer", "enable or disable developer mode",
			"gc", "garbage collect orphaned system extensions",
			"list", "list system extensions",
			"reset", "reset all system extensions state",
			"uninstall", "uninstall a system extension",
		),
	)
}
