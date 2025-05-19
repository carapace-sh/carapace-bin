package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Runs a defined package script",
	Aliases: []string{"run-script"},
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	runCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	runCmd.Flags().Bool("color", false, "Controls colors in the output")
	runCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	runCmd.Flags().String("filter", "", "set filter")
	runCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	runCmd.Flags().BoolP("help", "h", false, "Output usage information")
	runCmd.Flags().Bool("if-present", false, "Avoid exiting with a non-zero exit code when the script is undefined")
	runCmd.Flags().String("loglevel", "", "What level of logs to report")
	runCmd.Flags().Bool("no-bail", false, "The command will exit with a 0 exit code even if the script fails")
	runCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	runCmd.Flags().Bool("parallel", false, "Completely disregard concurrency and topological sorting")
	runCmd.Flags().BoolP("recursive", "r", false, "Run the defined package script in every package found in subdirectories")
	runCmd.Flags().Bool("report-summary", false, "Save the execution results of every package to \"pnpm-exec-summary.json\"")
	runCmd.Flags().Bool("resume-from", false, "Command executed from given package")
	runCmd.Flags().Bool("sequential", false, "Run the specified scripts one by one")
	runCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	runCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	runCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	runCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})

	// TODO complete scripts
}
