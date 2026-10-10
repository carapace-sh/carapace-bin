package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Run a task when files change",
	Aliases: []string{"w"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().StringP("glob", "g", "", "Glob pattern to watch")
	watchCmd.Flags().StringP("run", "r", "", "Task to run")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
