package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild a package",
	Aliases: []string{"rb"},
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()

	rebuildCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	rebuildCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	rebuildCmd.Flags().Bool("color", false, "Controls colors in the output")
	rebuildCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	rebuildCmd.Flags().String("filter", "", "set filter")
	rebuildCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	rebuildCmd.Flags().BoolP("help", "h", false, "Output usage information")
	rebuildCmd.Flags().String("loglevel", "", "What level of logs to report")
	rebuildCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	rebuildCmd.Flags().Bool("pending", false, "Rebuild packages that were not build during installation")
	rebuildCmd.Flags().BoolP("recursive", "r", false, "Rebuild every package found in subdirectories or every workspace package")
	rebuildCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	rebuildCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	rebuildCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	rebuildCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(rebuildCmd)

	carapace.Gen(rebuildCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})

	carapace.Gen(rebuildCmd).PositionalAnyCompletion(
		pnpm.ActionDependencies(),
	)
}
