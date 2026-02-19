package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a command to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_addCmd).Standalone()

	task_addCmd.Flags().String("arg", "", "Define an argument for the task")
	task_addCmd.Flags().Bool("clean-env", false, "Isolate the task from the shell environment")
	task_addCmd.Flags().String("cwd", "", "The working directory relative to the root of the project")
	task_addCmd.Flags().String("default-environment", "", "The default environment for the task")
	task_addCmd.Flags().String("depends-on", "", "Depends on these other commands")
	task_addCmd.Flags().String("description", "", "A short description of the task")
	task_addCmd.Flags().String("env", "", "A map of environment variables")
	task_addCmd.Flags().StringP("feature", "f", "", "The feature for which the task should be added")
	task_addCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	task_addCmd.Flags().StringP("platform", "p", "", "The platform for which the task should be added")
	taskCmd.AddCommand(task_addCmd)

	carapace.Gen(task_addCmd).FlagCompletion(carapace.ActionMap{
		"cwd":           carapace.ActionDirectories(),
		"feature":       pixi.ActionFeatures(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
