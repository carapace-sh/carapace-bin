package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Inspect tasks in concurrency groups. A same-named script takes precedence. Use `pnpm pm tasks` to force the built-in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_tasksCmd).Standalone()

	helpCmd.AddCommand(help_tasksCmd)
}
