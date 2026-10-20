// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var cache_taskCmd = &cobra.Command{
	Use:   "task",
	Short: "[experimental] Inspect artifact cache entries for a task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_taskCmd).Standalone()

	cache_taskCmd.Flags().BoolP("help", "h", false, "Print help")
	cache_taskCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	cacheCmd.AddCommand(cache_taskCmd)

	carapace.Gen(cache_taskCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
