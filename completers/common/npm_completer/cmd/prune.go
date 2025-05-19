package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove extraneous packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()
	pruneCmd.Flags().Bool("dry-run", false, "only report changes")
	pruneCmd.Flags().Bool("json", false, "output as json")
	pruneCmd.Flags().StringArray("omit", []string{""}, "omit package type")
	addWorkspaceFlags(pruneCmd)

	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		npm.ActionModules(),
	)
}
