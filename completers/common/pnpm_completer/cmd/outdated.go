package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:     "outdated",
	Short:   "Check for outdated packages",
	GroupID: "review",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	outdatedCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	outdatedCmd.Flags().Bool("color", false, "Controls colors in the output")
	outdatedCmd.Flags().Bool("compatible", false, "Print only versions that satisfy specs in package.json")
	outdatedCmd.Flags().BoolP("dev", "D", false, "Check only \"devDependencies\"")
	outdatedCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	outdatedCmd.Flags().String("filter", "", "set filter")
	outdatedCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	outdatedCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	outdatedCmd.Flags().BoolP("help", "h", false, "Output usage information")
	outdatedCmd.Flags().String("loglevel", "", "What level of logs to report")
	outdatedCmd.Flags().Bool("long", false, "By default, details about the outdated packages")
	outdatedCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	outdatedCmd.Flags().Bool("no-optional", false, "Don't check \"optionalDependencies\"")
	outdatedCmd.Flags().Bool("no-table", false, "Prints the outdated packages in a list")
	outdatedCmd.Flags().BoolP("prod", "P", false, "Check only \"dependencies\" and \"optionalDependencies\"")
	outdatedCmd.Flags().BoolP("recursive", "r", false, "Check for outdated dependencies in every package found in subdirectories")
	outdatedCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	outdatedCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	outdatedCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	outdatedCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		pnpm.ActionDependencies(),
	)
}
