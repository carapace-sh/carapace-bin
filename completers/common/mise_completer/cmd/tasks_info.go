// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var tasks_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show a task's definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_infoCmd).Standalone()

	tasks_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_infoCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	tasksCmd.AddCommand(tasks_infoCmd)

	carapace.Gen(tasks_infoCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
