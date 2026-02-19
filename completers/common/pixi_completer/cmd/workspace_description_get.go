package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_description_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_description_getCmd).Standalone()

	workspace_description_getCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_descriptionCmd.AddCommand(workspace_description_getCmd)

	carapace.Gen(workspace_description_getCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
