package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_dependencies_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add dependencies to the workspace dependencies table",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependencies_addCmd).Standalone()

	workspace_dependencies_addCmd.Flags().String("branch", "", "The git branch")
	workspace_dependencies_addCmd.Flags().StringP("git", "g", "", "The git URL for a workspace dependency")
	workspace_dependencies_addCmd.Flags().String("path", "", "The local path for a workspace dependency")
	workspace_dependencies_addCmd.Flags().String("rev", "", "The git revision")
	workspace_dependencies_addCmd.Flags().String("subdirectory", "", "The subdirectory of the git repository")
	workspace_dependencies_addCmd.Flags().String("tag", "", "The git tag")
	workspace_dependenciesCmd.AddCommand(workspace_dependencies_addCmd)

	carapace.Gen(workspace_dependencies_addCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})

	carapace.Gen(workspace_dependencies_addCmd).PositionalAnyCompletion(
		pixi.ActionPackageSearch(),
	)
}
