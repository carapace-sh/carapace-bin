package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/golangci-lint_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Lint the code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("allow-parallel-runners", false, "Allow multiple parallel golangci-lint instances running.")
	runCmd.Flags().Bool("allow-serial-runners", false, "Allow multiple golangci-lint instances running, but serialize them around a lock.")
	runCmd.Flags().StringSlice("build-tags", nil, "Build tags")
	runCmd.Flags().StringP("concurrency", "j", "", "Number of CPUs to use (Default: Automatically set to match Linux container CPU quota and fall back to the number of logical CPUs in the machine)")
	runCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	runCmd.PersistentFlags().String("cpu-profile-path", "", "Path to CPU profile output file")
	runCmd.Flags().String("default", "", "Default set of linters to enable")
	runCmd.Flags().StringSliceP("disable", "D", nil, "Disable specific linter")
	runCmd.Flags().StringSliceP("enable", "E", nil, "Enable specific linter")
	runCmd.Flags().StringSlice("enable-only", nil, "Override linters configuration section to only run the specific linter(s)")
	runCmd.Flags().Bool("fast-only", false, "Filter enabled linters to run only fast linters")
	runCmd.Flags().Bool("fix", false, "Fix found issues (if it's supported by the linter)")
	runCmd.Flags().Bool("internal-cmd-test", false, "Option is used only for testing golangci-lint command, don't use it")
	runCmd.Flags().String("issues-exit-code", "", "Exit code when issues were found")
	runCmd.Flags().String("max-issues-per-linter", "", "Maximum issues count per one linter. Set to 0 to disable")
	runCmd.Flags().String("max-same-issues", "", "Maximum count of issues with the same text. Set to 0 to disable")
	runCmd.PersistentFlags().String("mem-profile-path", "", "Path to memory profile output file")
	runCmd.Flags().String("modules-download-mode", "", "Modules download mode. If not empty, passed as -mod=<mode> to go tools")
	runCmd.Flags().BoolP("new", "n", false, "Show only new issues: if there are unstaged changes or untracked files, only those changes are analyzed, else only changes in HEAD~ are analyzed.")
	runCmd.Flags().String("new-from-merge-base", "", "Show only new issues created after the best common ancestor (merge-base against HEAD)")
	runCmd.Flags().String("new-from-patch", "", "Show only new issues created in git patch with file path `PATH`")
	runCmd.Flags().String("new-from-rev", "", "Show only new issues created after git revision `REV`")
	runCmd.Flags().Bool("no-config", false, "Don't read config file")
	runCmd.Flags().String("output.checkstyle.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().String("output.code-climate.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().String("output.html.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().String("output.json.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().Bool("output.junit-xml.extended", false, "Support extra JUnit XML fields.")
	runCmd.Flags().String("output.junit-xml.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().String("output.sarif.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().Bool("output.tab.colors", false, "Use colors.")
	runCmd.Flags().String("output.tab.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().Bool("output.tab.print-linter-name", false, "Print linter name in the end of issue text.")
	runCmd.Flags().String("output.teamcity.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().Bool("output.text.colors", false, "Use colors.")
	runCmd.Flags().String("output.text.path", "", "Output path can be either `stdout`, `stderr` or path to the file to write to.")
	runCmd.Flags().Bool("output.text.print-issued-lines", false, "Print lines of code with issue.")
	runCmd.Flags().Bool("output.text.print-linter-name", false, "Print linter name in the end of issue text.")
	runCmd.Flags().String("path-mode", "", "Path mode to use (empty, or 'abs')")
	runCmd.Flags().String("path-prefix", "", "Path prefix to add to output")
	runCmd.PersistentFlags().Bool("print-resources-usage", false, "Print avg and max memory usage of golangci-lint and total time")
	runCmd.Flags().Bool("show-stats", false, "Show statistics per linter")
	runCmd.Flags().StringSlice("skip-dirs", nil, "Regexps of directories to skip")
	runCmd.Flags().StringSlice("skip-files", nil, "Regexps of files to skip")
	runCmd.Flags().Bool("tests", false, "Analyze tests (*_test.go)")
	runCmd.Flags().String("timeout", "", "Timeout for total work. Disabled by default")
	runCmd.PersistentFlags().String("trace-path", "", "Path to trace output file")
	runCmd.Flags().Bool("uniq-by-line", false, "Make issues output unique by line")
	runCmd.Flags().Bool("whole-files", false, "Show issues in any part of update files (requires new-from-rev or new-from-patch)")
	runCmd.Flag("internal-cmd-test").Hidden = true
	runCmd.Flag("print-resources-usage").Hidden = true
	runCmd.Flag("skip-dirs").Hidden = true
	runCmd.Flag("skip-files").Hidden = true
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"build-tags":               golang.ActionBuildTags().UniqueList(","),
		"config":                   carapace.ActionFiles(),
		"cpu-profile-path":         carapace.ActionFiles(),
		"default":                  carapace.ActionValues("standard"), // TODO
		"disable":                  action.ActionLinters(runCmd).UniqueList(","),
		"enable":                   action.ActionLinters(runCmd).UniqueList(","),
		"enable-only":              action.ActionLinters(runCmd),
		"mem-profile-path":         carapace.ActionFiles(),
		"modules-download-mode":    golang.ActionModuleDownloadModes(),
		"new-from-merge-base":      git.ActionRefs(git.RefOption{}.Default()),
		"new-from-patch":           carapace.ActionFiles(),
		"new-from-rev":             git.ActionRefs(git.RefOption{}.Default()),
		"output.checkstyle.path":   action.ActionOutputPaths(),
		"output.code-climate.path": action.ActionOutputPaths(),
		"output.html.path":         action.ActionOutputPaths(),
		"output.json.path":         action.ActionOutputPaths(),
		"output.junit-xml.path":    action.ActionOutputPaths(),
		"output.sarif.path":        action.ActionOutputPaths(),
		"output.tab.path":          action.ActionOutputPaths(),
		"output.teamcity.path":     action.ActionOutputPaths(),
		"output.text.path":         action.ActionOutputPaths(),
		"path-mode":                carapace.ActionValues("empty", "abs"),
		"skip-dirs":                carapace.ActionDirectories().List(","),
		"skip-files":               carapace.ActionFiles().List(","),
		"trace-path":               carapace.ActionFiles(),
	})
}
