package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "uninstall", "un"},
	Short:   "Removes packages from `node_modules` and from the project's `package.json`",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	removeCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	removeCmd.Flags().Bool("color", false, "Controls colors in the output")
	removeCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	removeCmd.Flags().String("filter", "", "set filter")
	removeCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	removeCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	removeCmd.Flags().BoolP("help", "h", false, "Output usage information")
	removeCmd.Flags().String("loglevel", "", "What level of logs to report")
	removeCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	removeCmd.Flags().BoolP("recursive", "r", false, "Remove from every package found in subdirectories or from every workspace package")
	removeCmd.Flags().BoolP("save-dev", "D", false, "Remove the dependency only from \"devDependencies\"")
	removeCmd.Flags().BoolP("save-optional", "O", false, "Remove the dependency only from \"optionalDependencies\"")
	removeCmd.Flags().BoolP("save-prod", "P", false, "Remove the dependency only from \"dependencies\"")
	removeCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	removeCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	removeCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	removeCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		pnpm.ActionDependencies(),
	)
}
