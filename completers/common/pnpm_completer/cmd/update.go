package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up", "upgrade"},
	Short:   "Updates packages to their latest version based on the specified range",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	updateCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	updateCmd.Flags().Bool("color", false, "Controls colors in the output")
	updateCmd.Flags().String("depth", "", "How deep should levels of dependencies be inspected")
	updateCmd.Flags().BoolP("dev", "D", false, "Update packages only in \"devDependencies\"")
	updateCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	updateCmd.Flags().String("filter", "", "set filter")
	updateCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	updateCmd.Flags().BoolP("global", "g", false, "Update globally installed packages")
	updateCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	updateCmd.Flags().BoolP("help", "h", false, "Output usage information")
	updateCmd.Flags().BoolP("interactive", "i", false, "Show outdated dependencies and select which ones to update")
	updateCmd.Flags().BoolP("latest", "L", false, "Ignore version ranges in package.json")
	updateCmd.Flags().String("loglevel", "", "What level of logs to report")
	updateCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	updateCmd.Flags().Bool("no-optional", false, "Don't update packages in \"optionalDependencies\"")
	updateCmd.Flags().BoolP("prod", "P", false, "Update packages only in \"dependencies\" and \"optionalDependencies\"")
	updateCmd.Flags().BoolP("recursive", "r", false, "Update in every package found in subdirectories or every workspace package")
	updateCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	updateCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	updateCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	updateCmd.Flags().Bool("workspace", false, "Tries to link all packages from the workspace")
	updateCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
