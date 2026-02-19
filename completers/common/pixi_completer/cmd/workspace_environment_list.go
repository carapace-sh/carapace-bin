package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the environments in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_listCmd).Standalone()

	workspace_environment_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_environmentCmd.AddCommand(workspace_environment_listCmd)

	carapace.Gen(workspace_environment_listCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
