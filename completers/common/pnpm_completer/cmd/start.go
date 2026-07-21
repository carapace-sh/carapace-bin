package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Runs an arbitrary command specified in the package's start property of its scripts object",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	startCmd.Flags().String("filter", "", "set filter")
	startCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	startCmd.Flags().Bool("if-present", false, "Avoid exiting with a non-zero exit code when the script is undefined")
	startCmd.Flags().Bool("no-bail", false, "The command will exit with a 0 exit code even if the script fails")
	startCmd.Flags().Bool("parallel", false, "Completely disregard concurrency and topological sorting")
	startCmd.Flags().BoolP("recursive", "r", false, "Run the defined package script in every package found in subdirectories")
	startCmd.Flags().Bool("report-summary", false, "Save the execution results of every package to \"pnpm-exec-summary.json\"")
	startCmd.Flags().Bool("resume-from", false, "Command executed from given package")
	startCmd.Flags().Bool("sequential", false, "Run the specified scripts one by one")
	startCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
	})
}
