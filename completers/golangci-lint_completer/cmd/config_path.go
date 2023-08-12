package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golangcilint"
	"github.com/spf13/cobra"
)

var config_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print used config path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pathCmd).Standalone()

	config_pathCmd.Flags().Bool("allow-parallel-runners", false, "Allow multiple parallel golangci-lint instances running. If false (default) - golangci-lint acquires file lock on start.")
	config_pathCmd.Flags().Bool("allow-serial-runners", false, "Allow multiple golangci-lint instances running, but serialize them	around a lock. If false (default) - golangci-lint exits with an error if it fails to acquire file lock on start.")
	config_pathCmd.Flags().StringSlice("build-tags", []string{}, "Build tags")
	config_pathCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	config_pathCmd.Flags().String("deadline", "", "Deadline for total work")
	config_pathCmd.Flags().StringSliceP("disable", "D", []string{}, "Disable specific linter")
	config_pathCmd.Flags().Bool("disable-all", false, "Disable all linters")
	config_pathCmd.Flags().String("dupl.threshold", "", "Dupl: Minimal threshold to detect copy-paste")
	config_pathCmd.Flags().StringSliceP("enable", "E", []string{}, "Enable specific linter")
	config_pathCmd.Flags().Bool("enable-all", false, "Enable all linters")
	config_pathCmd.Flags().Bool("errcheck.check-blank", false, "Errcheck: check for errors assigned to blank identifier: _ = errFunc()")
	config_pathCmd.Flags().Bool("errcheck.check-type-assertions", false, "Errcheck: check for ignored type assertion results")
	config_pathCmd.Flags().String("errcheck.exclude", "", "Path to a file containing a list of functions to exclude from checking")
	config_pathCmd.Flags().String("errcheck.ignore", "", "Comma-separated list of pairs of the form pkg:regex. The regex is used to ignore names within pkg")
	config_pathCmd.Flags().StringSliceP("exclude", "e", []string{}, "Exclude issue by regexp")
	config_pathCmd.Flags().Bool("exclude-case-sensitive", false, "If set to true exclude and exclude rules regular expressions are case sensitive")
	config_pathCmd.Flags().Bool("exclude-use-default", false, "Use or not use default excludes:")
	config_pathCmd.Flags().Bool("fast", false, "Run only fast linters from enabled linters set (first run won't be fast)")
	config_pathCmd.Flags().Bool("fix", false, "Fix found issues (if it's supported by the linter)")
	config_pathCmd.Flags().String("go", "", "Targeted Go version")
	config_pathCmd.Flags().Bool("goconst.ignore-calls", false, "Goconst: ignore when constant is not used as function argument")
	config_pathCmd.Flags().Bool("goconst.match-constant", false, "Goconst: look for existing constants matching the values")
	config_pathCmd.Flags().String("goconst.max", "", "maximum value, only works with goconst.numbers")
	config_pathCmd.Flags().String("goconst.min", "", "minimum value, only works with goconst.numbers")
	config_pathCmd.Flags().String("goconst.min-len", "", "Goconst: minimum constant string length")
	config_pathCmd.Flags().String("goconst.min-occurrences", "", "Goconst: minimum occurrences of constant string count to trigger issue")
	config_pathCmd.Flags().Bool("goconst.numbers", false, "Goconst: search also for duplicated numbers")
	config_pathCmd.Flags().String("gocyclo.min-complexity", "", "Minimal complexity of function to report it")
	config_pathCmd.Flags().Bool("gofmt.simplify", false, "Gofmt: simplify code")
	config_pathCmd.Flags().String("golint.min-confidence", "", "Golint: minimum confidence of a problem to print it")
	config_pathCmd.Flags().Bool("govet.check-shadowing", false, "Govet: check for shadowed variables")
	config_pathCmd.Flags().Bool("internal-cmd-test", false, "Option is used only for testing golangci-lint command, don't use it")
	config_pathCmd.Flags().String("issues-exit-code", "", "Exit code when issues were found")
	config_pathCmd.Flags().String("lll.tab-width", "", "Lll: tab width in spaces")
	config_pathCmd.Flags().Bool("maligned.suggest-new", false, "Maligned: print suggested more optimal struct fields ordering")
	config_pathCmd.Flags().String("max-issues-per-linter", "", "Maximum issues count per one linter. Set to 0 to disable")
	config_pathCmd.Flags().String("max-same-issues", "", "Maximum count of issues with the same text. Set to 0 to disable")
	config_pathCmd.Flags().String("modules-download-mode", "", "Modules download mode. If not empty, passed as -mod=<mode> to go tools")
	config_pathCmd.Flags().BoolP("new", "n", false, "Show only new issues: if there are unstaged changes or untracked files, only those changes are analyzed, else only changes in HEAD~ are analyzed.")
	config_pathCmd.Flags().String("new-from-patch", "", "Show only new issues created in git patch with file path `PATH`")
	config_pathCmd.Flags().String("new-from-rev", "", "Show only new issues created after git revision `REV`")
	config_pathCmd.Flags().Bool("no-config", false, "Don't read config")
	config_pathCmd.Flags().String("out-format", "", "Format of output: colored-line-number|line-number|json|tab|checkstyle|code-climate|html|junit-xml|github-actions|teamcity")
	config_pathCmd.Flags().String("path-prefix", "", "Path prefix to add to output")
	config_pathCmd.Flags().StringSliceP("presets", "p", []string{}, "Enable presets (bugs|comment|complexity|error|format|import|metalinter|module|performance|sql|style|test|unused) of linters. Run 'golangci-lint linters' to see them. This option implies option --disable-all")
	config_pathCmd.Flags().Bool("print-issued-lines", false, "Print lines of code with issue")
	config_pathCmd.Flags().Bool("print-linter-name", false, "Print linter name in issue line")
	config_pathCmd.Flags().Bool("print-resources-usage", false, "Print avg and max memory usage of golangci-lint and total time")
	config_pathCmd.Flags().Bool("print-welcome", false, "Print welcome message")
	config_pathCmd.Flags().StringSlice("skip-dirs", []string{}, "Regexps of directories to skip")
	config_pathCmd.Flags().Bool("skip-dirs-use-default", false, "Use or not use default excluded directories:")
	config_pathCmd.Flags().StringSlice("skip-files", []string{}, "Regexps of files to skip")
	config_pathCmd.Flags().Bool("sort-results", false, "Sort linter results")
	config_pathCmd.Flags().Bool("tests", false, "Analyze tests (*_test.go)")
	config_pathCmd.Flags().String("timeout", "", "Timeout for total work")
	config_pathCmd.Flags().Bool("uniq-by-line", false, "Make issues output unique by line")
	config_pathCmd.Flags().Bool("whole-files", false, "Show issues in any part of update files (requires new-from-rev or new-from-patch)")
	config_pathCmd.Flag("deadline").Hidden = true
	config_pathCmd.Flag("dupl.threshold").Hidden = true
	config_pathCmd.Flag("errcheck.check-blank").Hidden = true
	config_pathCmd.Flag("errcheck.check-type-assertions").Hidden = true
	config_pathCmd.Flag("errcheck.exclude").Hidden = true
	config_pathCmd.Flag("errcheck.ignore").Hidden = true
	config_pathCmd.Flag("goconst.ignore-calls").Hidden = true
	config_pathCmd.Flag("goconst.match-constant").Hidden = true
	config_pathCmd.Flag("goconst.max").Hidden = true
	config_pathCmd.Flag("goconst.min").Hidden = true
	config_pathCmd.Flag("goconst.min-len").Hidden = true
	config_pathCmd.Flag("goconst.min-occurrences").Hidden = true
	config_pathCmd.Flag("goconst.numbers").Hidden = true
	config_pathCmd.Flag("gocyclo.min-complexity").Hidden = true
	config_pathCmd.Flag("gofmt.simplify").Hidden = true
	config_pathCmd.Flag("golint.min-confidence").Hidden = true
	config_pathCmd.Flag("govet.check-shadowing").Hidden = true
	config_pathCmd.Flag("internal-cmd-test").Hidden = true
	config_pathCmd.Flag("lll.tab-width").Hidden = true
	config_pathCmd.Flag("maligned.suggest-new").Hidden = true
	config_pathCmd.Flag("print-welcome").Hidden = true
	configCmd.AddCommand(config_pathCmd)

	carapace.Gen(config_pathCmd).FlagCompletion(carapace.ActionMap{
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
