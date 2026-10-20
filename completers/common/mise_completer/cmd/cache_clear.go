package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var cache_clearCmd = &cobra.Command{
	Use:     "clear",
	Short:   "Delete cached data",
	Aliases: []string{"c", "clean"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_clearCmd).Standalone()

	cache_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	cache_clearCmd.Flags().Bool("outdate", false, "Mark all cache files as old")
	cache_clearCmd.Flags().String("task", "", "[experimental] Clear task artifact cache entries for a task name or pattern")
	cache_clearCmd.Flag("outdate").Hidden = true
	cacheCmd.AddCommand(cache_clearCmd)

	carapace.Gen(cache_clearCmd).FlagCompletion(carapace.ActionMap{
		"task": mise.ActionTasks(),
	})

	carapace.Gen(cache_clearCmd).PositionalAnyCompletion(
		mise.ActionTools(),
	)
}
