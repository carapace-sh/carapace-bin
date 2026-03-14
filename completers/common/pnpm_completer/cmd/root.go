package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pnpm",
	Short: "Fast, disk space efficient package manager",
	Long:  "https://pnpm.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "manage", Title: "Manage Commands"},
		&cobra.Group{ID: "review", Title: "Review Commands"},
		&cobra.Group{ID: "run", Title: "Run Commands"},
		&cobra.Group{ID: "store", Title: "Store Commands"},
		&cobra.Group{ID: "other", Title: "Other Commands"},
	)

	// TODO verify flags
	rootCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	rootCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	rootCmd.Flags().Bool("color", false, "Controls colors in the output")
	rootCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	rootCmd.Flags().String("filter", "", "set filter")
	rootCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	rootCmd.Flags().BoolP("help", "h", false, "Output usage information")
	rootCmd.Flags().String("loglevel", "", "What level of logs to report")
	rootCmd.Flags().Bool("no-bail", false, "The command will exit with a 0 exit code even if the script fails")
	rootCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	rootCmd.Flags().Bool("parallel", false, "Completely disregard concurrency and topological sorting")
	rootCmd.Flags().BoolP("recursive", "r", false, "Run the command for each project in the workspace")
	rootCmd.Flags().Bool("report-summary", false, "Save the execution results of every package to \"pnpm-exec-summary.json\"")
	rootCmd.Flags().Bool("resume-from", false, "Command executed from given package")
	rootCmd.Flags().Bool("sequential", false, "Run the specified scripts one by one")
	rootCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	rootCmd.Flags().String("test-pattern", "", "Defines files related to tests.")
	rootCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().Bool("workspace-root", false, "Run the command on the root workspace project")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":         carapace.ActionDirectories(),
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
		"loglevel":    pnpm.ActionLoglevels(),
	})
}

// TODO verify these (now only used in start)
func addWorkspaceFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("workspace", "w", []string{""}, "Enable running a command in the context of the given workspace")
	cmd.Flags().Bool("workspaces", false, "Enable running a command in the context of all workspaces")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"workspace": carapace.Batch(
			pnpm.ActionWorkspaces(),
			pnpm.ActionWorkspaceDependencies(),
		).ToA(),
	})
}
