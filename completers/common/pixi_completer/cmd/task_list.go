package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_listCmd).Standalone()

	task_listCmd.Flags().StringP("environment", "e", "", "The environment to list tasks for")
	task_listCmd.Flags().Bool("json", false, "Output in JSON format")
	task_listCmd.Flags().Bool("machine-readable", false, "Output in machine-readable format")
	task_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	task_listCmd.Flags().BoolP("summary", "s", false, "Display a summary of the tasks")
	taskCmd.AddCommand(task_listCmd)

	carapace.Gen(task_listCmd).FlagCompletion(carapace.ActionMap{
		"environment":   pixi.ActionEnvironments(),
		"manifest-path": carapace.ActionFiles(),
	})
}
