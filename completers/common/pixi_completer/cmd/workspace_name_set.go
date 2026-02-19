package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_name_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_name_setCmd).Standalone()

	workspace_name_setCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_nameCmd.AddCommand(workspace_name_setCmd)

	carapace.Gen(workspace_name_setCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
