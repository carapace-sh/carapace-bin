package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all tasks in the workspace",
	Aliases: []string{"ls", "l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_listCmd).Standalone()

	task_listCmd.Flags().StringP("environment", "e", "", "The environment the list should be generated for. If not specified, the default environment is used")
	task_listCmd.Flags().Bool("json", false, "List as json instead of a tree If not specified, the default environment is used")
	task_listCmd.Flags().BoolP("summary", "s", false, "Tasks available for this machine per environment")
	taskCmd.AddCommand(task_listCmd)

	carapace.Gen(task_listCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
	})
}
