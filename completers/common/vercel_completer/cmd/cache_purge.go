package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge cache for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_purgeCmd).Standalone()

	cache_purgeCmd.Flags().String("project", "", "Project name or ID")
	cache_purgeCmd.Flags().String("type", "", "Type of cache to purge")
	cache_purgeCmd.Flags().Bool("yes", false, "Skip confirmation")

	cacheCmd.AddCommand(cache_purgeCmd)

	carapace.Gen(cache_purgeCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(cache_purgeCmd),
		"type":    carapace.ActionValues("purge", "all"),
	})
}
