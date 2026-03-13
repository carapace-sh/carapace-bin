package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Executes a shell command in scope of a project",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	execCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	execCmd.Flags().Bool("color", false, "Controls colors in the output")
	execCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	execCmd.Flags().String("filter", "", "set filter")
	execCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	execCmd.Flags().BoolP("help", "h", false, "Output usage information")
	execCmd.Flags().String("loglevel", "", "What level of logs to report")
	execCmd.Flags().Bool("no-bail", false, "The command will exit with a 0 exit code even if the script fails")
	execCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	execCmd.Flags().Bool("parallel", false, "Completely disregard concurrency and topological sorting")
	execCmd.Flags().BoolP("recursive", "r", false, "Run the defined package script in every package found in subdirectories")
	execCmd.Flags().Bool("report-summary", false, "Save the execution results of every package to \"pnpm-exec-summary.json\"")
	execCmd.Flags().Bool("resume-from", false, "Command executed from given package")
	execCmd.Flags().Bool("sequential", false, "Run the specified scripts one by one")
	execCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	execCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	execCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	execCmd.Flags().Bool("workspace-root", false, "Run the command on the root workspace project")
	addWorkspaceFlags(execCmd)
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
		"loglevel":    pnpm.ActionLoglevels(),
	})
}
