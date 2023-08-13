package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golangcilint"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the linters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("allow-parallel-runners", false, "Allow multiple parallel golangci-lint instances running. If false (default) - golangci-lint acquires file lock on start.")
	runCmd.Flags().Bool("allow-serial-runners", false, "Allow multiple golangci-lint instances running, but serialize them	around a lock. If false (default) - golangci-lint exits with an error if it fails to acquire file lock on start.")
	runCmd.Flags().StringSlice("build-tags", []string{}, "Build tags")
	runCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	runCmd.Flags().String("deadline", "", "Deadline for total work")
	runCmd.Flags().StringSliceP("disable", "D", []string{}, "Disable specific linter")
	runCmd.Flags().Bool("disable-all", false, "Disable all linters")
	runCmd.Flags().String("dupl.threshold", "", "Dupl: Minimal threshold to detect copy-paste")
	runCmd.Flags().StringSliceP("enable", "E", []string{}, "Enable specific linter")
	runCmd.Flags().Bool("enable-all", false, "Enable all linters")
	runCmd.Flags().Bool("errcheck.check-blank", false, "Errcheck: check for errors assigned to blank identifier: _ = errFunc()")
	runCmd.Flags().Bool("errcheck.check-type-assertions", false, "Errcheck: check for ignored type assertion results")
	runCmd.Flags().String("errcheck.exclude", "", "Path to a file containing a list of functions to exclude from checking")
	runCmd.Flags().String("errcheck.ignore", "", "Comma-separated list of pairs of the form pkg:regex. The regex is used to ignore names within pkg")
	runCmd.Flags().StringSliceP("exclude", "e", []string{}, "Exclude issue by regexp")
	runCmd.Flags().Bool("exclude-case-sensitive", false, "If set to true exclude and exclude rules regular expressions are case sensitive")
	runCmd.Flags().Bool("exclude-use-default", false, "Use or not use default excludes:")
	runCmd.Flags().Bool("fast", false, "Run only fast linters from enabled linters set (first run won't be fast)")
	runCmd.Flags().Bool("fix", false, "Fix found issues (if it's supported by the linter)")
	runCmd.Flags().String("go", "", "Targeted Go version")
	runCmd.Flags().Bool("goconst.ignore-calls", false, "Goconst: ignore when constant is not used as function argument")
	runCmd.Flags().Bool("goconst.match-constant", false, "Goconst: look for existing constants matching the values")
	runCmd.Flags().String("goconst.max", "", "maximum value, only works with goconst.numbers")
	runCmd.Flags().String("goconst.min", "", "minimum value, only works with goconst.numbers")
	runCmd.Flags().String("goconst.min-len", "", "Goconst: minimum constant string length")
	runCmd.Flags().String("goconst.min-occurrences", "", "Goconst: minimum occurrences of constant string count to trigger issue")
	runCmd.Flags().Bool("goconst.numbers", false, "Goconst: search also for duplicated numbers")
	runCmd.Flags().String("gocyclo.min-complexity", "", "Minimal complexity of function to report it")
	runCmd.Flags().Bool("gofmt.simplify", false, "Gofmt: simplify code")
	runCmd.Flags().String("golint.min-confidence", "", "Golint: minimum confidence of a problem to print it")
	runCmd.Flags().Bool("govet.check-shadowing", false, "Govet: check for shadowed variables")
	runCmd.Flags().Bool("internal-cmd-test", false, "Option is used only for testing golangci-lint command, don't use it")
	runCmd.Flags().String("issues-exit-code", "", "Exit code when issues were found")
	runCmd.Flags().String("lll.tab-width", "", "Lll: tab width in spaces")
	runCmd.Flags().Bool("maligned.suggest-new", false, "Maligned: print suggested more optimal struct fields ordering")
	runCmd.Flags().String("max-issues-per-linter", "", "Maximum issues count per one linter. Set to 0 to disable")
	runCmd.Flags().String("max-same-issues", "", "Maximum count of issues with the same text. Set to 0 to disable")
	runCmd.Flags().String("modules-download-mode", "", "Modules download mode. If not empty, passed as -mod=<mode> to go tools")
	runCmd.Flags().BoolP("new", "n", false, "Show only new issues: if there are unstaged changes or untracked files, only those changes are analyzed, else only changes in HEAD~ are analyzed.")
	runCmd.Flags().String("new-from-patch", "", "Show only new issues created in git patch with file path `PATH`")
	runCmd.Flags().String("new-from-rev", "", "Show only new issues created after git revision `REV`")
	runCmd.Flags().Bool("no-config", false, "Don't read config")
	runCmd.Flags().String("out-format", "", "Format of output: colored-line-number|line-number|json|tab|checkstyle|code-climate|html|junit-xml|github-actions|teamcity")
	runCmd.Flags().String("path-prefix", "", "Path prefix to add to output")
	runCmd.Flags().StringSliceP("presets", "p", []string{}, "Enable presets (bugs|comment|complexity|error|format|import|metalinter|module|performance|sql|style|test|unused) of linters. Run 'golangci-lint linters' to see them. This option implies option --disable-all")
	runCmd.Flags().Bool("print-issued-lines", false, "Print lines of code with issue")
	runCmd.Flags().Bool("print-linter-name", false, "Print linter name in issue line")
	runCmd.Flags().Bool("print-resources-usage", false, "Print avg and max memory usage of golangci-lint and total time")
	runCmd.Flags().Bool("print-welcome", false, "Print welcome message")
	runCmd.Flags().StringSlice("skip-dirs", []string{}, "Regexps of directories to skip")
	runCmd.Flags().Bool("skip-dirs-use-default", false, "Use or not use default excluded directories:")
	runCmd.Flags().StringSlice("skip-files", []string{}, "Regexps of files to skip")
	runCmd.Flags().Bool("sort-results", false, "Sort linter results")
	runCmd.Flags().Bool("tests", false, "Analyze tests (*_test.go)")
	runCmd.Flags().String("timeout", "", "Timeout for total work")
	runCmd.Flags().Bool("uniq-by-line", false, "Make issues output unique by line")
	runCmd.Flags().Bool("whole-files", false, "Show issues in any part of update files (requires new-from-rev or new-from-patch)")
	runCmd.Flag("deadline").Hidden = true
	runCmd.Flag("dupl.threshold").Hidden = true
	runCmd.Flag("errcheck.check-blank").Hidden = true
	runCmd.Flag("errcheck.check-type-assertions").Hidden = true
	runCmd.Flag("errcheck.exclude").Hidden = true
	runCmd.Flag("errcheck.ignore").Hidden = true
	runCmd.Flag("goconst.ignore-calls").Hidden = true
	runCmd.Flag("goconst.match-constant").Hidden = true
	runCmd.Flag("goconst.max").Hidden = true
	runCmd.Flag("goconst.min").Hidden = true
	runCmd.Flag("goconst.min-len").Hidden = true
	runCmd.Flag("goconst.min-occurrences").Hidden = true
	runCmd.Flag("goconst.numbers").Hidden = true
	runCmd.Flag("gocyclo.min-complexity").Hidden = true
	runCmd.Flag("gofmt.simplify").Hidden = true
	runCmd.Flag("golint.min-confidence").Hidden = true
	runCmd.Flag("govet.check-shadowing").Hidden = true
	runCmd.Flag("internal-cmd-test").Hidden = true
	runCmd.Flag("lll.tab-width").Hidden = true
	runCmd.Flag("maligned.suggest-new").Hidden = true
	runCmd.Flag("print-welcome").Hidden = true
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"build-tags":       golang.ActionBuildTags().UniqueList(","),
		"config":           carapace.ActionFiles(),
		"disable":          golangcilint.ActionLinters().UniqueList(","),
		"enable":           golangcilint.ActionLinters().UniqueList(","),
		"errcheck.exclude": carapace.ActionFiles(),
		"errcheck.ignore": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return golang.ActionPackages()
			default:
				return carapace.ActionValues()
			}
		}).List(","),
		"go":                    golang.ActionVersions(),
		"modules-download-mode": golang.ActionModuleDownloadModes(),
		"new-from-patch":        carapace.ActionFiles(),
		"new-from-rev":          git.ActionRefs(git.RefOption{}.Default()),
		"out-format":            golangcilint.ActionOutFormats(),
		"presets":               golangcilint.ActionPresets(),
		"skip-dirs":             carapace.ActionDirectories().List(","),
		"skip-files":            carapace.ActionFiles().List(","),
	})
}
