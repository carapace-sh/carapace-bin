package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "cleanup apps by removing old versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanupCmd).Standalone()
	cleanupCmd.Flags().BoolP("all", "a", false, "cleanup all apps (alternative to `*`)")
	cleanupCmd.Flags().BoolP("cache", "k", false, "remove outdated download cache")
	cleanupCmd.Flags().BoolP("global", "g", false, "cleanup a globally installed app")
	rootCmd.AddCommand(cleanupCmd)

	carapace.Gen(cleanupCmd).PositionalAnyCompletion(
		action.ActionInstalledApps(),
	)
}
