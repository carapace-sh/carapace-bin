package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias another specific command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_aliasCmd).Standalone()

	task_aliasCmd.Flags().String("description", "", "A short description of the alias")
	task_aliasCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	task_aliasCmd.Flags().StringP("platform", "p", "", "The platform for which the alias should be added")
	taskCmd.AddCommand(task_aliasCmd)

	carapace.Gen(task_aliasCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
