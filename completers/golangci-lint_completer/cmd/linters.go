package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golangcilint"
	"github.com/spf13/cobra"
)

var lintersCmd = &cobra.Command{
	Use:   "linters",
	Short: "List current linters configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintersCmd).Standalone()

	lintersCmd.Flags().Bool("allow-parallel-runners", false, "Allow multiple parallel golangci-lint instances running. If false (default) - golangci-lint acquires file lock on start.")
	lintersCmd.Flags().Bool("allow-serial-runners", false, "Allow multiple golangci-lint instances running, but serialize them	around a lock. If false (default) - golangci-lint exits with an error if it fails to acquire file lock on start.")
	lintersCmd.Flags().StringSlice("build-tags", []string{}, "Build tags")
	lintersCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	lintersCmd.Flags().String("deadline", "", "Deadline for total work")
	lintersCmd.Flags().StringSliceP("disable", "D", []string{}, "Disable specific linter")
	lintersCmd.Flags().Bool("disable-all", false, "Disable all linters")
	lintersCmd.Flags().String("dupl.threshold", "", "Dupl: Minimal threshold to detect copy-paste")
	lintersCmd.Flags().StringSliceP("enable", "E", []string{}, "Enable specific linter")
	lintersCmd.Flags().Bool("enable-all", false, "Enable all linters")
	lintersCmd.Flags().Bool("errcheck.check-blank", false, "Errcheck: check for errors assigned to blank identifier: _ = errFunc()")
	lintersCmd.Flags().Bool("errcheck.check-type-assertions", false, "Errcheck: check for ignored type assertion results")
	lintersCmd.Flags().String("errcheck.exclude", "", "Path to a file containing a list of functions to exclude from checking")
	lintersCmd.Flags().String("errcheck.ignore", "", "Comma-separated list of pairs of the form pkg:regex. The regex is used to ignore names within pkg")
	lintersCmd.Flags().StringSliceP("exclude", "e", []string{}, "Exclude issue by regexp")
	lintersCmd.Flags().Bool("exclude-case-sensitive", false, "If set to true exclude and exclude rules regular expressions are case sensitive")
	lintersCmd.Flags().Bool("exclude-use-default", false, "Use or not use default excludes:")
	lintersCmd.Flags().Bool("fast", false, "Run only fast linters from enabled linters set (first run won't be fast)")
	lintersCmd.Flags().Bool("fix", false, "Fix found issues (if it's supported by the linter)")
	lintersCmd.Flags().String("go", "", "Targeted Go version")
	lintersCmd.Flags().Bool("goconst.ignore-calls", false, "Goconst: ignore when constant is not used as function argument")
	lintersCmd.Flags().Bool("goconst.match-constant", false, "Goconst: look for existing constants matching the values")
	lintersCmd.Flags().String("goconst.max", "", "maximum value, only works with goconst.numbers")
	lintersCmd.Flags().String("goconst.min", "", "minimum value, only works with goconst.numbers")
	lintersCmd.Flags().String("goconst.min-len", "", "Goconst: minimum constant string length")
	lintersCmd.Flags().String("goconst.min-occurrences", "", "Goconst: minimum occurrences of constant string count to trigger issue")
	lintersCmd.Flags().Bool("goconst.numbers", false, "Goconst: search also for duplicated numbers")
	lintersCmd.Flags().String("gocyclo.min-complexity", "", "Minimal complexity of function to report it")
	lintersCmd.Flags().Bool("gofmt.simplify", false, "Gofmt: simplify code")
	lintersCmd.Flags().String("golint.min-confidence", "", "Golint: minimum confidence of a problem to print it")
	lintersCmd.Flags().Bool("govet.check-shadowing", false, "Govet: check for shadowed variables")
	lintersCmd.Flags().Bool("internal-cmd-test", false, "Option is used only for testing golangci-lint command, don't use it")
	lintersCmd.Flags().String("issues-exit-code", "", "Exit code when issues were found")
	lintersCmd.Flags().String("lll.tab-width", "", "Lll: tab width in spaces")
	lintersCmd.Flags().Bool("maligned.suggest-new", false, "Maligned: print suggested more optimal struct fields ordering")
	lintersCmd.Flags().String("max-issues-per-linter", "", "Maximum issues count per one linter. Set to 0 to disable")
	lintersCmd.Flags().String("max-same-issues", "", "Maximum count of issues with the same text. Set to 0 to disable")
	lintersCmd.Flags().String("modules-download-mode", "", "Modules download mode. If not empty, passed as -mod=<mode> to go tools")
	lintersCmd.Flags().BoolP("new", "n", false, "Show only new issues: if there are unstaged changes or untracked files, only those changes are analyzed, else only changes in HEAD~ are analyzed.")
	lintersCmd.Flags().String("new-from-patch", "", "Show only new issues created in git patch with file path `PATH`")
	lintersCmd.Flags().String("new-from-rev", "", "Show only new issues created after git revision `REV`")
	lintersCmd.Flags().Bool("no-config", false, "Don't read config")
	lintersCmd.Flags().String("out-format", "", "Format of output: colored-line-number|line-number|json|tab|checkstyle|code-climate|html|junit-xml|github-actions|teamcity")
	lintersCmd.Flags().String("path-prefix", "", "Path prefix to add to output")
	lintersCmd.Flags().StringSliceP("presets", "p", []string{}, "Enable presets (bugs|comment|complexity|error|format|import|metalinter|module|performance|sql|style|test|unused) of linters. Run 'golangci-lint linters' to see them. This option implies option --disable-all")
	lintersCmd.Flags().Bool("print-issued-lines", false, "Print lines of code with issue")
	lintersCmd.Flags().Bool("print-linter-name", false, "Print linter name in issue line")
	lintersCmd.Flags().Bool("print-resources-usage", false, "Print avg and max memory usage of golangci-lint and total time")
	lintersCmd.Flags().Bool("print-welcome", false, "Print welcome message")
	lintersCmd.Flags().StringSlice("skip-dirs", []string{}, "Regexps of directories to skip")
	lintersCmd.Flags().Bool("skip-dirs-use-default", false, "Use or not use default excluded directories:")
	lintersCmd.Flags().StringSlice("skip-files", []string{}, "Regexps of files to skip")
	lintersCmd.Flags().Bool("sort-results", false, "Sort linter results")
	lintersCmd.Flags().Bool("tests", false, "Analyze tests (*_test.go)")
	lintersCmd.Flags().String("timeout", "", "Timeout for total work")
	lintersCmd.Flags().Bool("uniq-by-line", false, "Make issues output unique by line")
	lintersCmd.Flags().Bool("whole-files", false, "Show issues in any part of update files (requires new-from-rev or new-from-patch)")
	lintersCmd.Flag("deadline").Hidden = true
	lintersCmd.Flag("dupl.threshold").Hidden = true
	lintersCmd.Flag("errcheck.check-blank").Hidden = true
	lintersCmd.Flag("errcheck.check-type-assertions").Hidden = true
	lintersCmd.Flag("errcheck.exclude").Hidden = true
	lintersCmd.Flag("errcheck.ignore").Hidden = true
	lintersCmd.Flag("goconst.ignore-calls").Hidden = true
	lintersCmd.Flag("goconst.match-constant").Hidden = true
	lintersCmd.Flag("goconst.max").Hidden = true
	lintersCmd.Flag("goconst.min").Hidden = true
	lintersCmd.Flag("goconst.min-len").Hidden = true
	lintersCmd.Flag("goconst.min-occurrences").Hidden = true
	lintersCmd.Flag("goconst.numbers").Hidden = true
	lintersCmd.Flag("gocyclo.min-complexity").Hidden = true
	lintersCmd.Flag("gofmt.simplify").Hidden = true
	lintersCmd.Flag("golint.min-confidence").Hidden = true
	lintersCmd.Flag("govet.check-shadowing").Hidden = true
	lintersCmd.Flag("internal-cmd-test").Hidden = true
	lintersCmd.Flag("lll.tab-width").Hidden = true
	lintersCmd.Flag("maligned.suggest-new").Hidden = true
	lintersCmd.Flag("print-welcome").Hidden = true
	rootCmd.AddCommand(lintersCmd)

	carapace.Gen(lintersCmd).FlagCompletion(carapace.ActionMap{
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
