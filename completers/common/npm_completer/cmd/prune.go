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
	pruneCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	pruneCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	pruneCmd.Flags().String("include", "", "include dependency types")
	pruneCmd.Flags().Bool("json", false, "output as json")
	pruneCmd.Flags().StringArray("omit", []string{""}, "omit package type")
	addWorkspaceFlags(pruneCmd)

	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).FlagCompletion(carapace.ActionMap{
		"include": carapace.ActionValues("prod", "dev", "optional", "peer"),
		"omit":    carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		npm.ActionModules(),
	)
}
