package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Commands to manage workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_descriptionCmd).Standalone()

	workspace_descriptionCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspaceCmd.AddCommand(workspace_descriptionCmd)

	carapace.Gen(workspace_descriptionCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
