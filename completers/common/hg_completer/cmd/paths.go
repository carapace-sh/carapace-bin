package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var pathsCmd = &cobra.Command{
	Use:     "paths",
	Short:   "show aliases for remote repositories",
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pathsCmd).Standalone()

	pathsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(pathsCmd)

	carapace.Gen(pathsCmd).PositionalCompletion(
		hg.ActionPaths(),
	)
}
