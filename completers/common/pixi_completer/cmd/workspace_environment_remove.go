package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environment from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_removeCmd).Standalone()

	workspace_environment_removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_environmentCmd.AddCommand(workspace_environment_removeCmd)

	carapace.Gen(workspace_environment_removeCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
