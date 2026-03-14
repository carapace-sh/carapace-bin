package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs an arbitrary command specified in the package's start property of its scripts object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	startCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	startCmd.Flags().Bool("color", false, "Controls colors in the output")
	startCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	startCmd.Flags().String("filter", "", "set filter")
	startCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	startCmd.Flags().BoolP("help", "h", false, "Output usage information")
	startCmd.Flags().Bool("if-present", false, "Avoid exiting with a non-zero exit code when the script is undefined")
	startCmd.Flags().String("loglevel", "", "What level of logs to report")
	startCmd.Flags().Bool("no-bail", false, "The command will exit with a 0 exit code even if the script fails")
	startCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	startCmd.Flags().Bool("parallel", false, "Completely disregard concurrency and topological sorting")
	startCmd.Flags().BoolP("recursive", "r", false, "Run the defined package script in every package found in subdirectories")
	startCmd.Flags().Bool("report-summary", false, "Save the execution results of every package to \"pnpm-exec-summary.json\"")
	startCmd.Flags().Bool("resume-from", false, "Command executed from given package")
	startCmd.Flags().Bool("sequential", false, "Run the specified scripts one by one")
	startCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	startCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	startCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	startCmd.Flags().Bool("workspace-root", false, "Run the command on the root workspace project")
	addWorkspaceFlags(startCmd)
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
		"loglevel":    pnpm.ActionLoglevels(),
	})
}
