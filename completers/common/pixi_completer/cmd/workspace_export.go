package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Commands to export workspaces to other formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_exportCmd).Standalone()

	workspace_exportCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspaceCmd.AddCommand(workspace_exportCmd)

	carapace.Gen(workspace_exportCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
