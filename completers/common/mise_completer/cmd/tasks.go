package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:     "tasks",
	Short:   "Manage tasks",
	Aliases: []string{"t"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasksCmd).Standalone()

	rootCmd.AddCommand(tasksCmd)

	carapace.Gen(tasksCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
